package main

import (
	"io"
	"os"

	"github.com/InfinityXOneSystems/logs"
)

func main() {
	// Example: Multi-output logging (console and file simultaneously)
	file, err := os.Create("combined.log")
	if err != nil {
		logs.Fatal("Failed to create log file")
	}
	defer file.Close()

	// Create a multi-writer that writes to both stdout and file
	multiWriter := io.MultiWriter(os.Stdout, file)

	logger := logs.New()
	logger.SetOutput(multiWriter)
	logger.SetLevel(logs.DEBUG)
	logger.SetPrefix("MultiApp")

	formatter := logs.NewTextFormatter()
	formatter.DisableColors = false // Colors for console (file will show codes)
	logger.SetFormatter(formatter)

	// Demonstrate various logging scenarios
	logger.Debug("Application initialization started")
	
	logger.InfoWithFields("Database connection established", logs.Fields{
		"host":     "localhost",
		"port":     5432,
		"database": "myapp",
	})

	logger.WarnWithFields("High memory usage detected", logs.Fields{
		"usage_mb":   512,
		"threshold":  400,
		"action":     "monitoring",
	})

	logger.ErrorWithFields("Failed to process request", logs.Fields{
		"request_id": "abc-123",
		"error_code": 500,
		"retry":      true,
	})

	logger.Info("Application running successfully")
	
	// Create a JSON logger for structured logs
	jsonLogger := logs.New()
	jsonFormatter := logs.NewJSONFormatter()
	jsonFormatter.PrettyPrint = true
	jsonLogger.SetFormatter(jsonFormatter)
	
	jsonLogger.InfoWithFields("JSON structured log", logs.Fields{
		"service":    "api",
		"version":    "1.0.0",
		"timestamp":  "2025-12-08T10:00:00Z",
		"metrics": map[string]int{
			"requests": 1000,
			"errors":   5,
		},
	})
}
