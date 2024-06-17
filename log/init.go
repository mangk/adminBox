package log

import (
	"os"
	"path"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/mangk/gAdmin/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _log *Log

func init() {
	// 日志基础配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.LevelKey = "_l"
	encoderConfig.TimeKey = "_t"
	encoderConfig.NameKey = "_n"
	encoderConfig.CallerKey = "_c"
	encoderConfig.StacktraceKey = "_s"

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
	if config.LogCfg().Prefix != "" {
		opt = append(opt, zap.Fields(zap.String("_p", config.LogCfg().Prefix)))
	}

	core := zapcore.NewCore(
		encoder,
		zapcore.NewMultiWriteSyncer(writer...),
		zap.NewAtomicLevelAt(zap.InfoLevel),
	)

	logger := zap.New(core, opt...)

	_log = &Log{
		Logger:     logger,
		callerSkip: 1,
	}
}

func Close() {
	_log.Logger.Sync()
}

func getLogfileWriter(dirName string) *rotatelogs.RotateLogs {
	maxAge := 30
	if config.LogCfg().MaxAge != 0 {
		maxAge = config.LogCfg().MaxAge
	}
	fileWriter, err := rotatelogs.New(
		path.Join(dirName, "%Y-%m-%d.log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(maxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic("设置日志输出错误:" + err.Error())
	}
	return fileWriter
}
