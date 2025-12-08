package logs

import (
	"encoding/json"
	"fmt"
	"time"
)

// Fields represents additional fields to include in a log entry
type Fields map[string]interface{}

// Entry represents a log entry
type Entry struct {
	Time    time.Time
	Level   Level
	Message string
	Fields  Fields
	Prefix  string
}

// Formatter is the interface for formatting log entries
type Formatter interface {
	Format(entry *Entry) ([]byte, error)
}

// TextFormatter formats log entries as human-readable text
type TextFormatter struct {
	TimestampFormat string
	DisableColors   bool
}

// NewTextFormatter creates a new text formatter with default settings
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		DisableColors:   false,
	}
}

// Format formats a log entry as text
func (f *TextFormatter) Format(entry *Entry) ([]byte, error) {
	timestamp := entry.Time.Format(f.TimestampFormat)
	level := entry.Level.String()

	if !f.DisableColors {
		level = colorizeLevel(entry.Level)
	}

	var msg string
	if entry.Prefix != "" {
		msg = fmt.Sprintf("[%s] %s [%s] %s", timestamp, level, entry.Prefix, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s] %s %s", timestamp, level, entry.Message)
	}

	if len(entry.Fields) > 0 {
		for k, v := range entry.Fields {
			msg += fmt.Sprintf(" %s=%v", k, v)
		}
	}

	msg += "\n"
	return []byte(msg), nil
}

// colorizeLevel adds ANSI color codes to log levels
func colorizeLevel(level Level) string {
	const (
		colorReset  = "\033[0m"
		colorRed    = "\033[31m"
		colorYellow = "\033[33m"
		colorBlue   = "\033[34m"
		colorGray   = "\033[90m"
	)

	switch level {
	case DEBUG:
		return colorGray + "DEBUG" + colorReset
	case INFO:
		return colorBlue + "INFO" + colorReset
	case WARN:
		return colorYellow + "WARN" + colorReset
	case ERROR:
		return colorRed + "ERROR" + colorReset
	case FATAL:
		return colorRed + "FATAL" + colorReset
	default:
		return level.String()
	}
}

// JSONFormatter formats log entries as JSON
type JSONFormatter struct {
	PrettyPrint bool
}

// NewJSONFormatter creates a new JSON formatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{
		PrettyPrint: false,
	}
}

// Format formats a log entry as JSON
func (f *JSONFormatter) Format(entry *Entry) ([]byte, error) {
	data := make(map[string]interface{})
	data["time"] = entry.Time.Format(time.RFC3339)
	data["level"] = entry.Level.String()
	data["message"] = entry.Message

	if entry.Prefix != "" {
		data["prefix"] = entry.Prefix
	}

	if len(entry.Fields) > 0 {
		for k, v := range entry.Fields {
			data[k] = v
		}
	}

	var output []byte
	var err error

	if f.PrettyPrint {
		output, err = json.MarshalIndent(data, "", "  ")
	} else {
		output, err = json.Marshal(data)
	}

	if err != nil {
		return nil, err
	}

	output = append(output, '\n')
	return output, nil
}
