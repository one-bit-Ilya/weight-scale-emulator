package config

import "github.com/one-bit-Ilya/weight-scale-emulator/internal/logger"

const (
	DefaultHost     string              = "127.0.0.1"
	DefaultPort     uint16              = 5001
	DefaultLogLevel logger.LoggingLevel = logger.LevelError
)
