package logger

// This class handles the logging configuration

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func SetLoggingLevel(loggingLevel string) {
	// set the global logging level
	switch loggingLevel {
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARNING":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL ":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "PANIC":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	}
}

func InitLogger(appEnv string, loggingLevel string) {
	SetLoggingLevel(loggingLevel)
	// decide on console logging vs JSON logging based on environment
	if appEnv != "DEV" {
		logger = zerolog.New(os.Stdout).With().Timestamp().Caller().
			//Str("service", serviceName).
			//Int("pid", os.Getpid()).
			Logger()
	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).
			With().
			Timestamp().
			Caller().
			//Str("service", serviceName).
			//Int("pid", os.Getpid()).
			Logger()
	}
}

func Info() *zerolog.Event {
	return logger.Info()
}

func Warn() *zerolog.Event {
	return logger.Warn()
}

func Debug() *zerolog.Event {
	return logger.Debug()
}

func Trace() *zerolog.Event {
	return logger.Trace()
}

func Fatal() *zerolog.Event {
	return logger.Fatal()
}

func Error() *zerolog.Event {
	return logger.Error()
}

func Err(err error) *zerolog.Event {
	return logger.Err(err)
}

func GetLevel() zerolog.Level {
	return logger.GetLevel()
}

func Hook(hook zerolog.Hook) zerolog.Logger {
	return logger.Hook(hook)
}

func Level(lvl zerolog.Level) zerolog.Logger {
	return logger.Level(lvl)
}

func Log(lvl zerolog.Level) *zerolog.Event {
	return logger.Log()
}

func Output(w io.Writer) zerolog.Logger {
	return logger.Output(w)
}

func Panic() *zerolog.Event {
	return logger.Panic()
}

func GetLoggerInstance() zerolog.Logger {
	return logger
}

// below is another formatting style

/* return zerolog.New(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	NoColor:    false,
	TimeFormat: time.StampMicro,
	FormatLevel: func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("[%s]", i))
	},
	PartsOrder: []string{
		zerolog.TimestampFieldName,
		zerolog.LevelFieldName,
		zerolog.CallerFieldName,
		zerolog.MessageFieldName,
	},
}).With().Timestamp().Caller().Logger() */
