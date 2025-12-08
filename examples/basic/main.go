package main

import (
	"os"

	"github.com/InfinityXOneSystems/logs"
)

func main() {
	// Example 1: Using the default logger
	logs.Info("Application started")
	logs.Debug("This is a debug message")
	logs.Warn("This is a warning")
	logs.Error("This is an error")

	// Example 2: Formatted logging
	logs.Infof("User %s logged in at %s", "john_doe", "2025-12-08 10:00:00")

	// Example 3: Logging with fields
	logs.InfoWithFields("User action", logs.Fields{
		"user":   "john_doe",
		"action": "login",
		"ip":     "192.168.1.1",
	})

	// Example 4: Custom logger with JSON formatter
	jsonLogger := logs.New()
	jsonLogger.SetFormatter(logs.NewJSONFormatter())
	jsonLogger.Info("This is logged as JSON")
	jsonLogger.InfoWithFields("User action in JSON", logs.Fields{
		"user": "jane_doe",
		"role": "admin",
	})

	// Example 5: Custom logger with file output
	file, err := os.Create("app.log")
	if err != nil {
		logs.Error("Failed to create log file")
	} else {
		defer file.Close()
		fileLogger := logs.New()
		fileLogger.SetOutput(file)
		fileLogger.SetLevel(logs.DEBUG)
		fileLogger.Info("This is logged to a file")
		fileLogger.Debug("Debug information in file")
	}

	// Example 6: Logger with prefix
	logs.SetPrefix("MyApp")
	logs.Info("Application with prefix")

	// Example 7: Setting log level
	logs.SetLevel(logs.WARN)
	logs.Debug("This won't be shown")
	logs.Info("This won't be shown either")
	logs.Warn("This will be shown")
	logs.Error("This will also be shown")

	logs.Info("Application finished")
}
