package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mangk/adminBox/pkg/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _log *LogInstance
var _logInitMutex sync.Mutex

// Init initializes the global logger. It should be called once from the main application.
func init() {

	// 日志基础配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "_l"
	encoderConfig.TimeKey = "_t"
	encoderConfig.NameKey = "_n"
	encoderConfig.CallerKey = "_c"
	encoderConfig.StacktraceKey = "_s"
	encoderConfig.FunctionKey = "_f"

	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if config.LogCfg().Format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// 日志输出位置
	writerMap := make(map[string]zapcore.WriteSyncer)
	for _, output := range config.LogCfg().Output {
		if output == "console" {
			writerMap[output] = zapcore.AddSync(os.Stdout)
		} else {
			writerMap[output] = zapcore.AddSync(getLogfileWriter(output))
		}
	}
	var writer []zapcore.WriteSyncer
	for _, w := range writerMap {
		writer = append(writer, w)
	}
	if len(writer) == 0 {
		writer = append(writer, zapcore.AddSync(os.Stdout))
	}

	// 其他配置
	var opt []zap.Option
	opt = append(opt, zap.WithCaller(true), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(writer...),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logger := zap.New(core, opt...)

	_log = &LogInstance{
		logger:     logger,
		callerSkip: 1,
	}
}

func i() *LogInstance {
	if _log == nil {
		// This will panic if Init() has not been called.
		// This is intentional to enforce proper initialization.
		panic("logger has not been initialized. Please call log.Init() in your main function.")
	}
	return _log
}

func Info(msg string, keysAndValues ...interface{}) {
	i().Info(msg, keysAndValues...)
}

func Infof(format string, args ...interface{}) {
	i().Infof(format, args...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	i().Warn(msg, keysAndValues...)
}

func Warnf(format string, args ...interface{}) {
	i().Warnf(format, args...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	i().Debug(msg, keysAndValues...)
}

func Debugf(format string, args ...interface{}) {
	i().Debugf(format, args...)
}

func Error(msg string, keysAndValues ...interface{}) {
	i().Error(errors.New(msg), msg, keysAndValues...)
}

func Errorf(format string, args ...interface{}) {
	i().Errorf(format, args...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	i().Panic(msg, keysAndValues...)
}

func Panicf(format string, args ...interface{}) {
	i().Panicf(format, args...)
}

func JsonMarshal(data any) {
	if v, ok := data.(string); ok {
		i().Info(v)
		return
	}
	v, e := json.Marshal(data)
	i().Infof("%s; %e", v, e)
}

func Sync() {
	i().Sync()
}

func Print(args ...interface{}) {
	if len(args) > 0 {
		format := []string{}
		for i := 0; i < len(args); i++ {
			format = append(format, "%+v")
		}
		i().Infof(strings.Join(format, "\n"), args...)
	}
}

func Zaplog() *zap.Logger {
	return i().logger
}

func Trace(traceKey ...string) *LogInstance {
	if len(traceKey) == 0 {
		traceKey = append(traceKey, uuid.New().String())
	}
	return &LogInstance{traceKey: traceKey[0], callerSkip: 1, logger: i().logger}
}

// LogInstance
type LogInstance struct {
	callerSkip int
	traceKey   string
	logger     *zap.Logger
}

func (l *LogInstance) sugared() *zap.SugaredLogger {
	_l := l.logger.Sugar()
	if l.traceKey != "" {
		_l = _l.With("_trace", l.traceKey)
	}
	return _l.WithOptions(zap.AddCallerSkip(l.callerSkip))
}

func (l *LogInstance) Info(msg string, keysAndValues ...interface{}) {
	l.sugared().Infow(msg, keysAndValues...)
}

func (l *LogInstance) Infof(format string, args ...interface{}) {
	l.sugared().Infof(format, args...)
}

func (l *LogInstance) Warn(msg string, keysAndValues ...interface{}) {
	l.sugared().Warnw(msg, keysAndValues...)
}

func (l *LogInstance) Warnf(format string, args ...interface{}) {
	l.sugared().Warnf(format, args...)
}

func (l *LogInstance) Debug(msg string, keysAndValues ...interface{}) {
	l.sugared().Debugw(msg, keysAndValues...)
}

func (l *LogInstance) Debugf(format string, args ...interface{}) {
	l.sugared().Debugf(format, args...)
}

func (l *LogInstance) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "err", err)
	l.sugared().Errorw(msg, keysAndValues...)
}

func (l *LogInstance) Errorf(format string, args ...interface{}) {
	l.sugared().Errorf(format, args...)
}

func (l *LogInstance) Panic(msg string, keysAndValues ...interface{}) {
	l.sugared().Panicw(msg, keysAndValues...)
}

func (l *LogInstance) Panicf(format string, args ...interface{}) {
	l.sugared().Panicf(format, args...)
}

func (l *LogInstance) Sync() {
	l.sugared().Sync()
}

// GormLogger
type GormLogger struct {
	logger *LogInstance
}

func GormAdapter() *GormLogger {
	return &GormLogger{
		logger: i(),
	}
}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	// TODO 这里处理的不够全面，不够深入 gorm
	if len(args) == 4 {
		kv := []interface{}{}
		kv = append(kv, "call", args[0], "cost", fmt.Sprintf("%.3fms", args[1]), "rows", args[2], "sql", args[3])
		g.logger.sugared().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infow("_gorm", kv...)
	} else {
		g.logger.sugared().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infof(format, args...)
	}
}

// GinLogger
type GinLogger struct {
	logger *LogInstance
}

func GinAdapter() *GinLogger {
	return &GinLogger{
		logger: i(),
	}
}
func (g *GinLogger) Write(p []byte) (n int, err error) {
	g.logger.sugared().WithOptions(zap.WithCaller(false)).Infof("%s", p)
	return
}

func getLogfileWriter(dirName string) *rotatelogs.RotateLogs {
	maxAge := 30
	if config.LogCfg().MaxAge != 0 {
		maxAge = config.LogCfg().MaxAge
	}

	fileWriter, err := rotatelogs.New(
		path.Join(dirName, config.LogCfg().Prefix+"%Y-%m-%d.log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic("设置日志输出错误:" + err.Error())
	}
	return fileWriter
}

func Close() {
	i().logger.Sync()
}
