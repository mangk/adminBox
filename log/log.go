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
	"github.com/mangk/adminBox/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _log *logInstance
var _logInitOnce sync.Once

func Info(msg string, keysAndValues ...interface{}) {
	_log.Info(msg, keysAndValues...)
}

func Infof(format string, args ...interface{}) {
	_log.Infof(format, args...)
}

func Warn(msg string, keysAndValues ...interface{}) {
	_log.Warn(msg, keysAndValues...)
}

func Warnf(format string, args ...interface{}) {
	_log.Warnf(format, args...)
}

func Debug(msg string, keysAndValues ...interface{}) {
	_log.Debug(msg, keysAndValues...)
}

func Debugf(format string, args ...interface{}) {
	_log.Debugf(format, args...)
}

func Error(msg string, keysAndValues ...interface{}) {
	_log.Error(errors.New(msg), msg, keysAndValues...)
}

func Errorf(format string, args ...interface{}) {
	_log.Errorf(format, args...)
}

func Panic(msg string, keysAndValues ...interface{}) {
	_log.Panic(msg, keysAndValues...)
}

func Panicf(format string, args ...interface{}) {
	_log.Panicf(format, args...)
}

func Print(args ...interface{}) {
	if len(args) > 0 {
		format := []string{}
		for i := 0; i < len(args); i++ {
			format = append(format, "%+v")
		}
		_log.Infof(strings.Join(format, "\n"), args...)
	}
}

func Zaplog() *zap.Logger {
	return _log.logger
}

func Trace(traceKey ...string) *logInstance {
	if len(traceKey) == 0 {
		traceKey = append(traceKey, uuid.New().String())
	}
	return &logInstance{traceKey: traceKey[0], callerSkip: 1}
}

// logInstance
type logInstance struct {
	callerSkip int
	traceKey   string
	logger     *zap.Logger
}

func (l *logInstance) i() *zap.SugaredLogger {
	_logInitOnce.Do(func() {
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
		writer := []zapcore.WriteSyncer{}
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

		_log = &logInstance{
			logger:     logger,
			callerSkip: 1,
		}
	})

	_l := _log.logger.Sugar()
	if l != nil && l.traceKey != "" {
		_l = _l.With("_trace", l.traceKey)
		return _l
	}
	return _l.WithOptions(zap.AddCallerSkip(_log.callerSkip))
}

func (l *logInstance) Info(msg string, keysAndValues ...interface{}) {
	l.i().Infow(msg, keysAndValues...)
}

func (l *logInstance) Infof(format string, args ...interface{}) {
	l.i().Infof(format, args...)
}

func (l *logInstance) Warn(msg string, keysAndValues ...interface{}) {
	l.i().Warnw(msg, keysAndValues...)
}

func (l *logInstance) Warnf(format string, args ...interface{}) {
	l.i().Warnf(format, args...)
}

func (l *logInstance) Debug(msg string, keysAndValues ...interface{}) {
	l.i().Debugw(msg, keysAndValues...)
}

func (l *logInstance) Debugf(format string, args ...interface{}) {
	l.i().Debugf(format, args...)
}

func (l *logInstance) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "err", err)
	l.i().Errorw(msg, keysAndValues...)
}

func (l *logInstance) Errorf(format string, args ...interface{}) {
	l.i().Errorf(format, args...)
}

func (l *logInstance) Panic(msg string, keysAndValues ...interface{}) {
	l.i().Panicw(msg, keysAndValues...)
}

func (l *logInstance) Panicf(format string, args ...interface{}) {
	l.i().Panicf(format, args...)
}

// GormLogger
type GormLogger struct {
	logger *logInstance
}

func GormAdapter() *GormLogger {
	return &GormLogger{
		logger: _log,
	}
}

func (g *GormLogger) Printf(format string, args ...interface{}) {
	// TODO 这里处理的不够全面，不够深入 gorm
	if len(args) == 4 {
		kv := []interface{}{}
		kv = append(kv, "call", args[0], "cost", fmt.Sprintf("%.3fms", args[1]), "rows", args[2], "sql", args[3])
		g.logger.i().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infow("_gorm", kv...)
	} else {
		g.logger.i().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infof(format, args...)
	}
}

// GinLogger
type GinLogger struct {
	logger *logInstance
}

func GinAdapter() *GinLogger {
	return &GinLogger{
		logger: _log,
	}
}
func (g *GinLogger) Write(p []byte) (n int, err error) {
	// TODO 这里处理的不够全面，不够深入 gin
	args := []interface{}{}
	if e := json.Unmarshal(p, &args); e == nil {
		g.logger.i().WithOptions(zap.WithCaller(false)).Infow("httpRequest", args...)
	} else {
		g.logger.i().WithOptions(zap.AddCallerSkip(5)).Infof("%s", p)
	}
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
	_log.logger.Sync()
}
