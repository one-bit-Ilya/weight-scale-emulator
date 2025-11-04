# weight-scale-emulator

A configurable TCP server emulator for testing integration with industrial weighing equipment protocols.

## Overview

This emulator allows developers to test their integration code without requiring physical hardware.

## Current Implementation Status

### Implemented Features
- Configuration management via YAML file
- Structured logging with different levels (INFO, WARN, ERROR)
- Colored console output for better readability
- Default values for missing configuration
- Structures prepared for parsing requests
- Structures prepared for generating responses

### Planned Features
- TCP server implementation
- Protocol command handler
- CLI interface for manual control

## Quick Start

### Prerequisites
- Go 1.25.1 or later

### Installation

1. Clone the repository:
```bash
git clone https://github.com/one-bit-Ilya/weight-scale-emulator.git
cd weight-scale-emulator
```

2. Create configuration file `config.yaml`:
```yaml
server:
  host: "127.0.0.1"
  port: 5001
logging:
  level: "ERROR"
```

3. Run the emulator:
```bash
go run ./cmd/scale-emulator
```

## Configuration

The emulator uses config.yaml for configuration. If the file is missing or invalid, default values will be used:

|Setting | Default | Description |
|--------|---------|-------------|
|server.host | "127.0.0.1" | Server host address |
|server.port | 5001 | Server port |
| logging.level | "ERROR" | Logging level (INFO, WARN, ERROR) |

## Project Structure

```
weight-scale-emulator/
├── cmd
│   └── scale-emulator
│       └── main.go
├── config.yaml
├── go.mod
├── go.sum
├── internal
│   ├── cli
│   │   ├── cli.go
│   │   └── colors.go
│   ├── config
│   │   ├── config.go
│   │   └── defaults.go
│   ├── handler
│   ├── logger
│   │   └── logger.go
│   ├── protocol
│   │   ├── codec.go
│   │   ├── const.go
│   │   ├── crc.go
│   │   ├── request.go
│   │   └── response.go
│   ├── router
│   │   └── router.go
│   └── server
│       └── server.go
└── README.md
```