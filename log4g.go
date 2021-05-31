package log4g

import (
	"errors"
	"fmt"
	"time"
)

const (
	LOG4G_VERSION = "log4g-v0.99.0"
	LOG4G_MAJOR   = 0
	LOG4G_MINOR   = 99
	LOG4G_PATCH   = 0
)

var logLevel = INFO

type Logger struct {
	Date     time.Time
	Category string
	Level    Level
	Message  string
}

func SetLogLevel(level Level) (bool, error) {
	if int(level) > len(levelStrings) || level < 0 {
		return false, errors.New("Invalid log level")
	}

	logLevel = level
	return true, nil
}

func Category(category string) *Logger {
	logger := Logger{Category: category}

	return &logger
}

func (l *Logger) write() {
	if int(logLevel) <= int(l.Level) {
		fmt.Println(fmt.Sprintf("[%s] [%s] %-8s - %s", l.Date.Format(time.RFC3339), l.Category, l.Category, l.Message))
	}
}

func (l *Logger) handle(level Level, message string) {
	l.Level = level
	l.Message = message
	l.write()
}

func (l *Logger) Debug(message string) {
	l.handle(DEBUG, message)
}

func (l *Logger) Info(message string) {
	l.handle(INFO, message)
}

func (l *Logger) Warning(message string) {
	l.handle(WARNING, message)
}

func (l *Logger) Error(message string) {
	l.handle(ERROR, message)
}

func (l *Logger) Critical(message string) {
	l.handle(CRITICAL, message)
}

func (l *Logger) Fatal(message string) {
	l.handle(FATAL, message)
}
