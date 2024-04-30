package logger

import (
	"os"

	"github.com/rs/zerolog"
)

// New returns a configured zerolog.Logger based on the environment.
func New(environment string) zerolog.Logger {
	switch environment {
	case "development":
		zerolog.FormattedLevels = map[zerolog.Level]string{
			zerolog.DebugLevel: "DEBUG",
			zerolog.InfoLevel:  "INFO",
			zerolog.WarnLevel:  "WARN",
			zerolog.ErrorLevel: "ERROR",
			zerolog.FatalLevel: "FATAL",
			zerolog.PanicLevel: "PANIC",
		}

		output := zerolog.ConsoleWriter{Out: os.Stderr}

		return zerolog.New(output).
			Level(zerolog.DebugLevel).
			With().
			Timestamp().
			Logger()
	case "production":
		return zerolog.New(os.Stderr).
			Level(zerolog.ErrorLevel).
			With().
			Timestamp().
			Logger()
	default:
		return zerolog.New(os.Stderr).
			Level(zerolog.InfoLevel).
			With().
			Timestamp().
			Logger()
	}
}
