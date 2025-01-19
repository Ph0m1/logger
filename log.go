package logging

import (
	"errors"
	"io"
	"log"
	"os"
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

const (
	LevelDebug = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelCritical
)

var (
	ErrInvalidLogLevel = errors.New("invalid log level")
	defaultLogger      = BaseLogger{level: LevelCritical, logger: log.New(os.Stderr, "", log.LstdFlags)}
	logLevels          = map[string]int{
		"DEBUG":    LevelDebug,
		"INFO":     LevelInfo,
		"WARN":     LevelWarn,
		"ERROR":    LevelError,
		"CRITICAL": LevelCritical,
		"FATAL":    LevelFatal,
	}
)

type BaseLogger struct {
	level  int
	Prefix string
	logger *log.Logger
}

func NewLogger(level string, out io.Writer, prefix string) (BaseLogger, error) {
	l, ok := logLevels[level]
	if !ok {
		return defaultLogger, ErrInvalidLogLevel
	}
	return BaseLogger{level: l, Prefix: prefix, logger: log.New(out, "", log.LstdFlags)}, nil
}

func (l *BaseLogger) Debug(args ...interface{}) {
	if l.level > LevelDebug {
		return
	}
	l.outputLog("DEBUG:", args...)
}
func (l *BaseLogger) Info(args ...interface{}) {
	if l.level > LevelInfo {
		return
	}
	l.outputLog("INFO:", args...)
}

func (l *BaseLogger) Warn(args ...interface{}) {
	if l.level > LevelWarn {
		return
	}
	l.outputLog("WARN:", args...)
}

func (l *BaseLogger) Error(args ...interface{}) {
	if l.level > LevelError {
		return
	}
	l.outputLog("ERROR:", args...)
}

func (l *BaseLogger) Fatal(args ...interface{}) {
	if l.level > LevelFatal {
		return
	}
	l.outputLog("FATAL:", args...)
	os.Exit(1)
}

func (l *BaseLogger) outputLog(level string, args ...interface{}) {
	msg := make([]interface{}, len(args))
	msg[0] = l.Prefix
	msg[1] = level
	copy(msg[2:], args)
	l.logger.Println(msg...)
}
