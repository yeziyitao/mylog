package mylog

import (
	"github.com/go-mesh/openlogging"
)

// Debug Debug(action string, data ...openlogging.Option)
func Debug(action string, data ...openlogging.Option) {
	Logger.Debug(action, data...)
}

// Info Info(action string, data ...openlogging.Option)
func Info(action string, data ...openlogging.Option) {
	Logger.Info(action, data...)
}

// Warn Warn(action string, data ...openlogging.Option)
func Warn(action string, data ...openlogging.Option) {
	Logger.Warn(action, data...)
}

// Error Error(action string, data ...openlogging.Option)
func Error(action string, data ...openlogging.Option) {
	Logger.Error(action, data...)
}

// Fatal Fatal(action string, data ...openlogging.Option)
func Fatal(action string, data ...openlogging.Option) {
	Logger.Fatal(action, data...)
}

// Debugf Debugf(format string, args ...interface{})
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Infof Infof(format string, args ...interface{})
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warnf Warnf(format string, args ...interface{})
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Errorf Errorf(format string, args ...interface{})
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Fatalf Fatalf(format string, args ...interface{})
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}
