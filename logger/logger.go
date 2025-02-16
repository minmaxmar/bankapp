package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var appLogger zerolog.Logger

func InitLogger(logLevel string) {
	level := strings.ToLower(logLevel)
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	appLogger = zerolog.New(consoleWriter).With().Timestamp().Logger() //.Level(level)

	log.Logger = appLogger
}

func GetLogger() *zerolog.Logger {
	return &log.Logger
}
