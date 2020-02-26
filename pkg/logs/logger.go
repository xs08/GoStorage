package logs

import (
	"io"
	"os"

	"github.com/rs/zerolog"
)

// Logger log
type Logger struct {
	*zerolog.Logger
}

// global setting
// log.Logger
// zerolog.SetGlobalLevel
// zerolog.DisableSampling
// zerolog.TimestampFieldName
// zerolog.LevelFieldName
// zerolog.MessageFieldName
// zerolog.ErrorFieldName
// zerolog.TimeFieldFormat: zerolog.TimeFormatUnix, zerolog.TimeFormatUnixMs, zerolog.TimeFormatUnixMicro
// zerolog.DurationFieldUnit
// zerolog.DurationFieldInteger
// zerolog.ErrorHandler

// NewLogger get new logger
func NewLogger(out io.Writer, info map[string]string) *Logger {
	if out == os.Stdout {
		out = zerolog.ConsoleWriter{Out: os.Stderr}
	}
	loggerContext := zerolog.
		New(out).
		With().Timestamp()

	// extra infos
	if info != nil {
		for key, value := range info {
			loggerContext = loggerContext.Str(key, value)
		}
	}
	logger := loggerContext.Logger()

	logger.Debug().
		Msg("logger initialized")

	return &Logger{&logger}
}

func init() {
	// log time format
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// log level
	if os.Getenv("env") == "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	zerolog.TimestampFieldName = "_t"
	zerolog.LevelFieldName = "_level"
	zerolog.MessageFieldName = "_msg"
	zerolog.ErrorFieldName = "_err"
}
