# Quick Start Guide

## Installation

```bash
go get github.com/InfinityXOneSystems/logs
```

## 30-Second Example

```go
package main

import "github.com/InfinityXOneSystems/logs"

func main() {
    logs.Info("Hello, World!")
    logs.Warn("Something might be wrong")
    logs.Error("An error occurred")
}
```

## Common Use Cases

### 1. Simple Application Logging

```go
logs.Info("Server started on port 8080")
logs.Infof("Connected to database: %s", dbName)
```

### 2. Structured Logging for Analysis

```go
logs.InfoWithFields("User login", logs.Fields{
    "user_id": 12345,
    "ip": "192.168.1.1",
    "success": true,
})
```

### 3. JSON Logs for Production

```go
logger := logs.New()
logger.SetFormatter(logs.NewJSONFormatter())
logger.SetLevel(logs.INFO)
logger.Info("Production ready")
```

### 4. File Logging

```go
file, _ := os.Create("app.log")
defer file.Close()

logger := logs.New()
logger.SetOutput(file)
logger.Info("Logged to file")
```

### 5. Component-Specific Logging

```go
// Create loggers for different components
dbLogger := logs.New()
dbLogger.SetPrefix("DB")

apiLogger := logs.New()
apiLogger.SetPrefix("API")

dbLogger.Info("Connected")     // [timestamp] INFO [DB] Connected
apiLogger.Info("Request received") // [timestamp] INFO [API] Request received
```

## Log Levels (from most to least verbose)

- `DEBUG` - Detailed debugging information
- `INFO` - General information about application operation
- `WARN` - Warning messages for concerning situations
- `ERROR` - Error messages for failures
- `FATAL` - Critical errors that terminate the application

## Next Steps

- Check out the [README](README.md) for complete documentation
- Explore the [examples](examples/basic/main.go) for more use cases
- Run tests with `go test -v`
