package logs

import (
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

// Level represents the severity level of a log message
type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
	FATAL
)

// String returns the string representation of the log level
func (l Level) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARN"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return "UNKNOWN"
	}
}

// Logger represents a logger instance
type Logger struct {
	mu        sync.Mutex
	output    io.Writer
	level     Level
	formatter Formatter
	prefix    string
}

// New creates a new logger with the default configuration
func New() *Logger {
	return &Logger{
		output:    os.Stdout,
		level:     INFO,
		formatter: NewTextFormatter(),
		prefix:    "",
	}
}

// SetOutput sets the output destination for the logger
func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.output = w
}

// SetLevel sets the minimum log level
func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// SetFormatter sets the log formatter
func (l *Logger) SetFormatter(f Formatter) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.formatter = f
}

// SetPrefix sets the logger prefix
func (l *Logger) SetPrefix(prefix string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.prefix = prefix
}

// log writes a log message with the given level
func (l *Logger) log(level Level, msg string, fields Fields) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if level < l.level {
		return
	}

	entry := &Entry{
		Time:    time.Now(),
		Level:   level,
		Message: msg,
		Fields:  fields,
		Prefix:  l.prefix,
	}

	formatted, err := l.formatter.Format(entry)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting log: %v\n", err)
		return
	}

	_, err = l.output.Write(formatted)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing log: %v\n", err)
	}

	if level == FATAL {
		os.Exit(1)
	}
}

// Debug logs a debug message
func (l *Logger) Debug(msg string) {
	l.log(DEBUG, msg, nil)
}

// Debugf logs a formatted debug message
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.log(DEBUG, fmt.Sprintf(format, args...), nil)
}

// DebugWithFields logs a debug message with fields
func (l *Logger) DebugWithFields(msg string, fields Fields) {
	l.log(DEBUG, msg, fields)
}

// Info logs an info message
func (l *Logger) Info(msg string) {
	l.log(INFO, msg, nil)
}

// Infof logs a formatted info message
func (l *Logger) Infof(format string, args ...interface{}) {
	l.log(INFO, fmt.Sprintf(format, args...), nil)
}

// InfoWithFields logs an info message with fields
func (l *Logger) InfoWithFields(msg string, fields Fields) {
	l.log(INFO, msg, fields)
}

// Warn logs a warning message
func (l *Logger) Warn(msg string) {
	l.log(WARN, msg, nil)
}

// Warnf logs a formatted warning message
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.log(WARN, fmt.Sprintf(format, args...), nil)
}

// WarnWithFields logs a warning message with fields
func (l *Logger) WarnWithFields(msg string, fields Fields) {
	l.log(WARN, msg, fields)
}

// Error logs an error message
func (l *Logger) Error(msg string) {
	l.log(ERROR, msg, nil)
}

// Errorf logs a formatted error message
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.log(ERROR, fmt.Sprintf(format, args...), nil)
}

// ErrorWithFields logs an error message with fields
func (l *Logger) ErrorWithFields(msg string, fields Fields) {
	l.log(ERROR, msg, fields)
}

// Fatal logs a fatal message and exits
func (l *Logger) Fatal(msg string) {
	l.log(FATAL, msg, nil)
}

// Fatalf logs a formatted fatal message and exits
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.log(FATAL, fmt.Sprintf(format, args...), nil)
}

// FatalWithFields logs a fatal message with fields and exits
func (l *Logger) FatalWithFields(msg string, fields Fields) {
	l.log(FATAL, msg, fields)
}

// Default logger instance
var defaultLogger = New()

// Debug logs a debug message using the default logger
func Debug(msg string) {
	defaultLogger.Debug(msg)
}

// Debugf logs a formatted debug message using the default logger
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// DebugWithFields logs a debug message with fields using the default logger
func DebugWithFields(msg string, fields Fields) {
	defaultLogger.DebugWithFields(msg, fields)
}

// Info logs an info message using the default logger
func Info(msg string) {
	defaultLogger.Info(msg)
}

// Infof logs a formatted info message using the default logger
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// InfoWithFields logs an info message with fields using the default logger
func InfoWithFields(msg string, fields Fields) {
	defaultLogger.InfoWithFields(msg, fields)
}

// Warn logs a warning message using the default logger
func Warn(msg string) {
	defaultLogger.Warn(msg)
}

// Warnf logs a formatted warning message using the default logger
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// WarnWithFields logs a warning message with fields using the default logger
func WarnWithFields(msg string, fields Fields) {
	defaultLogger.WarnWithFields(msg, fields)
}

// Error logs an error message using the default logger
func Error(msg string) {
	defaultLogger.Error(msg)
}

// Errorf logs a formatted error message using the default logger
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// ErrorWithFields logs an error message with fields using the default logger
func ErrorWithFields(msg string, fields Fields) {
	defaultLogger.ErrorWithFields(msg, fields)
}

// Fatal logs a fatal message and exits using the default logger
func Fatal(msg string) {
	defaultLogger.Fatal(msg)
}

// Fatalf logs a formatted fatal message and exits using the default logger
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// FatalWithFields logs a fatal message with fields and exits using the default logger
func FatalWithFields(msg string, fields Fields) {
	defaultLogger.FatalWithFields(msg, fields)
}

// SetOutput sets the output destination for the default logger
func SetOutput(w io.Writer) {
	defaultLogger.SetOutput(w)
}

// SetLevel sets the minimum log level for the default logger
func SetLevel(level Level) {
	defaultLogger.SetLevel(level)
}

// SetFormatter sets the log formatter for the default logger
func SetFormatter(f Formatter) {
	defaultLogger.SetFormatter(f)
}

// SetPrefix sets the logger prefix for the default logger
func SetPrefix(prefix string) {
	defaultLogger.SetPrefix(prefix)
}
