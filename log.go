/*
 * Copyright 2024 hiifong <i@hiif.ong>
 */

package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	defaultLoggerOnce sync.Once
	defaultLogger     *zap.Logger
)

// Default default logger
func Default() *zap.Logger {
	defaultLoggerOnce.Do(func() {
		if defaultLogger == nil {
			ec := zap.NewDevelopmentEncoderConfig()
			ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
			enc := zapcore.NewConsoleEncoder(ec)
			core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
			options := []zap.Option{
				zap.AddStacktrace(zapcore.ErrorLevel),
				zap.AddCaller(),
				zap.AddCallerSkip(1),
			}
			defaultLogger = zap.New(core, options...)
		}
	})

	return defaultLogger
}

// SetDefault set a default logger
func SetDefault(logger *zap.Logger) {
	defaultLogger = logger
}

// Logger return a default logger
func Logger() *zap.Logger {
	return defaultLogger
}

// Debug ...
func Debug(msg string, fields ...zap.Field) {
	Default().Debug(msg, fields...)
}

// Info ...
func Info(msg string, fields ...zap.Field) {
	Default().Info(msg, fields...)
}

// Warn ...
func Warn(msg string, fields ...zap.Field) {
	Default().Warn(msg, fields...)
}

// Error ...
func Error(msg string, fields ...zap.Field) {
	Default().Error(msg, fields...)
}

// Panic ...
func Panic(msg string, fields ...zap.Field) {
	Default().Panic(msg, fields...)
}

// Debugf ...
func Debugf(format string, args ...interface{}) {
	Default().Sugar().Debugf(format, args...)
}

// Infof ...
func Infof(format string, args ...interface{}) {
	Default().Sugar().Infof(format, args...)
}

// Warnf ...
func Warnf(format string, args ...interface{}) {
	Default().Sugar().Warnf(format, args...)
}

// Errorf ...
func Errorf(format string, args ...interface{}) {
	Default().Sugar().Errorf(format, args...)
}

// Fatalf ...
func Fatalf(format string, args ...interface{}) {
	Default().Sugar().Fatalf(format, args...)
}

// Sync call zap logger Sync
func Sync() {
	err := defaultLogger.Sync()
	if err != nil {
		Error(err.Error())
	}
	Info("logger has been synced")
}
