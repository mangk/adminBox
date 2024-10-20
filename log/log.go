package log

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

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
		format := ""
		for i := 0; i < len(args); i++ {
			format += "%+v\n"
		}
		_log.Infof(format, args...)
	}
}

func Logger() *Log {
	return _log
}

func LoggerAdapter(name string) *Log {
	return &Log{
		CallerSkip: 0,
		traceKey:   name,
		Logger:     _log.Logger,
	}
}

func Zaplog() *zap.Logger {
	return _log.Logger
}

func Trace(traceKey ...string) *Log {
	if len(traceKey) == 0 {
		traceKey = append(traceKey, uuid.New().String())
	}
	return &Log{traceKey: traceKey[0]}
}

// Log
type Log struct {
	CallerSkip int
	traceKey   string
	Logger     *zap.Logger
}

func (l *Log) SugaredLogger() *zap.SugaredLogger {
	_l := _log.Logger.Sugar()
	if l.traceKey != "" {
		_l = _l.With("_trace", l.traceKey)
	}
	return _l.WithOptions(zap.AddCallerSkip(l.CallerSkip))
}

func (l *Log) Info(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger().Infow(msg, keysAndValues...)
}

func (l *Log) Infof(format string, args ...interface{}) {
	l.SugaredLogger().Infof(format, args...)
}

func (l *Log) Warn(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger().Warnw(msg, keysAndValues...)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	l.SugaredLogger().Warnf(format, args...)
}

func (l *Log) Debug(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger().Debugw(msg, keysAndValues...)
}

func (l *Log) Debugf(format string, args ...interface{}) {
	l.SugaredLogger().Debugf(format, args...)
}

func (l *Log) Error(err error, msg string, keysAndValues ...interface{}) {
	keysAndValues = append(keysAndValues, "err", err)
	l.SugaredLogger().Errorw(msg, keysAndValues...)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	l.SugaredLogger().Errorf(format, args...)
}

func (l *Log) Panic(msg string, keysAndValues ...interface{}) {
	l.SugaredLogger().Panicw(msg, keysAndValues...)
}

func (l *Log) Panicf(format string, args ...interface{}) {
	l.SugaredLogger().Panicf(format, args...)
}

// GormLogger
type GormLogger struct {
	logger *Log
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
		g.logger.SugaredLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infow("_gorm", kv...)
	} else {
		g.logger.SugaredLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infof(format, args...)
	}
}

// GinLogger
type GinLogger struct {
	logger *Log
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
		g.logger.SugaredLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infow("_gin", args...)
	} else {
		g.logger.SugaredLogger().WithOptions(zap.AddCallerSkip(1), zap.WithCaller(false)).Infof("%s", p)
	}
	return
}
