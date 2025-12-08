# logs

A simple, flexible, and efficient logging library for Go applications.

## Features

- **Multiple Log Levels**: DEBUG, INFO, WARN, ERROR, and FATAL
- **Flexible Formatters**: Text and JSON formatters included
- **Customizable Output**: Write to console, files, or any `io.Writer`
- **Structured Logging**: Support for additional fields
- **Thread-Safe**: Safe for concurrent use
- **Zero Dependencies**: Uses only Go standard library
- **Colorized Output**: Optional colored output for better readability
- **Prefixes**: Add prefixes to organize logs from different components

## Installation

```bash
go get github.com/InfinityXOneSystems/logs
```

## Quick Start

```go
package main

import "github.com/InfinityXOneSystems/logs"

func main() {
    logs.Info("Application started")
    logs.Warn("This is a warning")
    logs.Error("This is an error")
}
```

## Usage

### Basic Logging

```go
logs.Debug("Debug message")
logs.Info("Info message")
logs.Warn("Warning message")
logs.Error("Error message")
logs.Fatal("Fatal message") // Exits with status 1
```

### Formatted Logging

```go
logs.Infof("User %s logged in", username)
logs.Errorf("Failed to connect to %s:%d", host, port)
```

### Structured Logging with Fields

```go
logs.InfoWithFields("User action", logs.Fields{
    "user":   "john_doe",
    "action": "login",
    "ip":     "192.168.1.1",
})
```

### Custom Logger

```go
logger := logs.New()
logger.SetLevel(logs.DEBUG)
logger.SetPrefix("MyApp")
logger.Info("Custom logger message")
```

### JSON Formatter

```go
logger := logs.New()
logger.SetFormatter(logs.NewJSONFormatter())
logger.Info("This will be logged as JSON")
```

### File Output

```go
file, err := os.Create("app.log")
if err != nil {
    panic(err)
}
defer file.Close()

logger := logs.New()
logger.SetOutput(file)
logger.Info("This is logged to a file")
```

### Setting Log Level

```go
logs.SetLevel(logs.WARN) // Only WARN, ERROR, and FATAL will be logged
```

## Log Levels

The following log levels are available (in order of severity):

- `DEBUG`: Detailed information for debugging
- `INFO`: General informational messages
- `WARN`: Warning messages for potentially harmful situations
- `ERROR`: Error messages for serious issues
- `FATAL`: Critical errors that cause the application to exit

## Formatters

### Text Formatter (Default)

Outputs logs in a human-readable format:
```
[2025-12-08 10:00:00] INFO Application started
[2025-12-08 10:00:01] WARN This is a warning
```

### JSON Formatter

Outputs logs in JSON format for structured logging:
```json
{"time":"2025-12-08T10:00:00Z","level":"INFO","message":"Application started"}
```

## Examples

See the [examples](examples/) directory for more usage examples.

## API Reference

### Logger Methods

- `New()` - Create a new logger instance
- `SetOutput(w io.Writer)` - Set the output destination
- `SetLevel(level Level)` - Set the minimum log level
- `SetFormatter(f Formatter)` - Set the log formatter
- `SetPrefix(prefix string)` - Set a prefix for log messages

### Logging Methods

Each log level has three methods:
- Basic: `Info(msg string)`
- Formatted: `Infof(format string, args ...interface{})`
- With Fields: `InfoWithFields(msg string, fields Fields)`

Available for: Debug, Info, Warn, Error, Fatal

## Testing

Run the tests with:

```bash
go test -v
```

## License

MIT License - see LICENSE file for details