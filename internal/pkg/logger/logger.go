// Package logger logger base on zap
//
// The differents between development logger add production logger
// is that development logger include caller.
// Thus, a production log should like
// {"level":"ERROR","time":"2019-07-18T03:09:46.098+0800","message":"cert and key can not be empty, failed to listen https port"}
// a development log should like
// {"level":"ERROR","time":"2019-07-18T03:12:45.599+0800","caller":"logger/logger.go:129","message":"cert and key can not be empty, failed to listen https port"}
//
// Development logger output through log file and console
// Production logger output just log file
package logger

import (
	"fmt"
	"os"

	"github.com/puti-projects/puti/internal/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

// InitLogger init zap logger
func InitLogger(runmode string) {
	if runmode == "release" {
		logger = InitProductionLogger()
	} else {
		logger = InitDevelopmentLogger()
	}
	defer logger.Sync()

	zap.ReplaceGlobals(logger)
}

// InitProductionLogger init the logger for production environment
func InitProductionLogger() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(newPutiEncoderConfig()),
		getWriteSyncer("error"),
		highPriority,
	)

	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

// InitDevelopmentLogger init the logger for development environment
func InitDevelopmentLogger() *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	core := zapcore.NewTee(
		// normal json encoder
		zapcore.NewCore(
			zapcore.NewJSONEncoder(newPutiEncoderConfig()),
			getWriteSyncer("info"),
			lowPriority,
		),
		// error json encoder
		zapcore.NewCore(
			zapcore.NewJSONEncoder(newPutiEncoderConfig()),
			getWriteSyncer("error"),
			highPriority,
		),
		// normal console
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(newPutiEncoderConfig()),
			zapcore.Lock(os.Stdout),
			lowPriority,
		),
		// error console
		zapcore.NewCore(
			zapcore.NewConsoleEncoder(newPutiEncoderConfig()),
			zapcore.Lock(os.Stderr),
			highPriority,
		),
	)
	return zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func getWriteSyncer(writemode string) zapcore.WriteSyncer {
	sl := &lumberjack.Logger{
		MaxSize:    config.Log.LoggerMaxSize, // megabytes
		MaxBackups: config.Log.LoggerMaxBackups,
		MaxAge:     config.Log.LoggerMaxAge, // days
	}

	if writemode == "error" {
		sl.Filename = config.Log.LoggerFileError
		return zapcore.AddSync(sl)
	}

	sl.Filename = config.Log.LoggerFileInfo
	return zapcore.AddSync(sl)
}

func newPutiEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05"),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log site.
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	logger.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log site.
func Warn(msg string, args ...zapcore.Field) {
	logger.Warn(msg, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	logger.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at ErrorLevel. The message includes any fields passed at the log site.
func Error(msg string, args ...zapcore.Field) {
	logger.Error(msg, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	logger.Error(fmt.Sprintf(template, args...))
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed at the log site.
// DPanic means "development panic"
func DPanic(msg string, args ...zapcore.Field) {
	logger.DPanic(msg, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message.
// In development, the logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	logger.DPanic(fmt.Sprintf(template, args...))
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log site.
func Panic(msg string, args ...zapcore.Field) {
	logger.Panic(msg, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	logger.Panic(fmt.Sprintf(template, args...))
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log site.
func Fatal(msg string, args ...zapcore.Field) {
	logger.Fatal(msg, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	logger.Fatal(fmt.Sprintf(template, args...))
}
