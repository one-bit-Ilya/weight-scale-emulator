package logger

import (
	"log"
	"os"

	"github.com/one-bit-Ilya/weight-scale-emulator/internal/cli"
)

type LoggingLevel string

const (
	LevelInfo  LoggingLevel = "INFO"
	LevelWarn  LoggingLevel = "WARN"
	LevelError LoggingLevel = "ERROR"
)

func IsValidLoggingLevel(level LoggingLevel) bool {
	switch level {
	case LevelInfo, LevelWarn, LevelError:
		return true
	}
	return false
}

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func Init() {
	InfoLog = log.New(os.Stdout, cli.Blue+"[INFO]: "+cli.Reset, log.Ldate|log.Ltime)
	WarningLog = log.New(os.Stdout, cli.Yellow+"[WARN]: "+cli.Reset, log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLog = log.New(os.Stdout, cli.Red+"[ERROR]: "+cli.Reset, log.Ldate|log.Ltime|log.Lshortfile)
}
