package logs

import (
	"bytes"
	"strings"
	"testing"
)

func TestLoggerLevels(t *testing.T) {
	tests := []struct {
		name     string
		level    Level
		expected string
	}{
		{"DEBUG", DEBUG, "DEBUG"},
		{"INFO", INFO, "INFO"},
		{"WARN", WARN, "WARN"},
		{"ERROR", ERROR, "ERROR"},
		{"FATAL", FATAL, "FATAL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.level.String() != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, tt.level.String())
			}
		})
	}
}

func TestLoggerOutput(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(DEBUG)
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	logger.SetFormatter(formatter)

	logger.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "INFO") {
		t.Errorf("Expected output to contain INFO, got: %s", output)
	}
	if !strings.Contains(output, "test message") {
		t.Errorf("Expected output to contain 'test message', got: %s", output)
	}
}

func TestLoggerWithFields(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(DEBUG)
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	logger.SetFormatter(formatter)

	fields := Fields{
		"user": "testuser",
		"id":   123,
	}
	logger.InfoWithFields("user action", fields)

	output := buf.String()
	if !strings.Contains(output, "user=testuser") {
		t.Errorf("Expected output to contain 'user=testuser', got: %s", output)
	}
	if !strings.Contains(output, "id=123") {
		t.Errorf("Expected output to contain 'id=123', got: %s", output)
	}
}

func TestLoggerPrefix(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(DEBUG)
	logger.SetPrefix("MyApp")
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	logger.SetFormatter(formatter)

	logger.Info("test message")

	output := buf.String()
	if !strings.Contains(output, "[MyApp]") {
		t.Errorf("Expected output to contain '[MyApp]', got: %s", output)
	}
}

func TestLoggerLevelFiltering(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(WARN)
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	logger.SetFormatter(formatter)

	logger.Debug("debug message")
	logger.Info("info message")

	output := buf.String()
	if strings.Contains(output, "debug message") {
		t.Errorf("Debug message should not be logged when level is WARN")
	}
	if strings.Contains(output, "info message") {
		t.Errorf("Info message should not be logged when level is WARN")
	}

	logger.Warn("warn message")
	output = buf.String()
	if !strings.Contains(output, "warn message") {
		t.Errorf("Warn message should be logged when level is WARN")
	}
}

func TestLoggerFormatted(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(DEBUG)
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	logger.SetFormatter(formatter)

	logger.Infof("user %s logged in with id %d", "john", 42)

	output := buf.String()
	if !strings.Contains(output, "user john logged in with id 42") {
		t.Errorf("Expected formatted message, got: %s", output)
	}
}

func TestJSONFormatter(t *testing.T) {
	logger := New()
	buf := &bytes.Buffer{}
	logger.SetOutput(buf)
	logger.SetLevel(DEBUG)
	logger.SetFormatter(NewJSONFormatter())

	logger.Info("json test")

	output := buf.String()
	if !strings.Contains(output, `"level":"INFO"`) {
		t.Errorf("Expected JSON output to contain level, got: %s", output)
	}
	if !strings.Contains(output, `"message":"json test"`) {
		t.Errorf("Expected JSON output to contain message, got: %s", output)
	}
}

func TestDefaultLogger(t *testing.T) {
	buf := &bytes.Buffer{}
	SetOutput(buf)
	SetLevel(DEBUG)
	
	formatter := NewTextFormatter()
	formatter.DisableColors = true
	SetFormatter(formatter)

	Info("default logger test")

	output := buf.String()
	if !strings.Contains(output, "default logger test") {
		t.Errorf("Expected output from default logger, got: %s", output)
	}
}
