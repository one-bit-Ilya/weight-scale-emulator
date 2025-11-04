package config

import (
	"os"

	"github.com/one-bit-Ilya/weight-scale-emulator/internal/logger"
	"gopkg.in/yaml.v3"
)

type ServerSetting struct {
	Host string `yaml:"host"`
	Port uint16 `yaml:"port"`
}

type LoggingLevelSetting struct {
	LoggingLevel logger.LoggingLevel `yaml:"level"`
}

type Config struct {
	Server  ServerSetting       `yaml:"server"`
	Logging LoggingLevelSetting `yaml:"logging"`
}

func LoadConfig() *Config {
	var cfg Config
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		cfg.Server.Host = DefaultHost
		cfg.Server.Port = DefaultPort
		cfg.Logging.LoggingLevel = logger.LevelError
		logger.WarningLog.Println("Failed to read configuration file - using default values")
		return &cfg
	}
	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		cfg.Server.Host = DefaultHost
		cfg.Server.Port = DefaultPort
		cfg.Logging.LoggingLevel = logger.LevelError
		logger.WarningLog.Println("Failed to unmarshal configuration file - using default values")
		return &cfg
	}
	if cfg.Server.Host == "" {
		cfg.Server.Host = DefaultHost
		logger.WarningLog.Printf("Failed to set host - using default value (%v)\n", DefaultHost)
	}
	if cfg.Server.Port == 0 {
		cfg.Server.Port = DefaultPort
		logger.WarningLog.Printf("Failed to set port - using default value (%v)\n", DefaultPort)
	}
	if !logger.IsValidLoggingLevel(cfg.Logging.LoggingLevel) {
		cfg.Logging.LoggingLevel = DefaultLogLevel
		logger.WarningLog.Printf("Failed to set logging level - using default value (%v)\n", DefaultLogLevel)
	}
	return &cfg
}
