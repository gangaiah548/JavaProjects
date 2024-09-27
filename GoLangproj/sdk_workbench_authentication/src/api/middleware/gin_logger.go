package middleware

/*

 Reference : https://learninggolang.com/it5-gin-structured-logging.html

*/

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// DefaultStructuredLogger logs a gin HTTP request in JSON format. Uses the
// default logger from rs/zerolog.
func DefaultStructuredLogger(env string) gin.HandlerFunc {
	// change logging style based on environment
	var logger zerolog.Logger
	if env != "DEV" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()
	} else {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}).With().Timestamp().Caller().Logger()
	}

	return StructuredLogger(logger)
}

// StructuredLogger logs a gin HTTP request in JSON format. Allows to set the
// logger for testing purposes.
func StructuredLogger(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// Log using the params
		var logEvent *zerolog.Event
		if c.Writer.Status() >= 500 {
			logEvent = logger.Error()
		} else {
			logEvent = logger.Info()
		}

		logEvent.Str("client_id", param.ClientIP).
			Str("method", param.Method).
			Int("status_code", param.StatusCode).
			Int("body_size", param.BodySize).
			Str("path", param.Path).
			Str("latency", param.Latency.String()).
			Msg(param.ErrorMessage)
	}
}

// Recovery is a Gin middleware that recovers both panics and requests with errors by logging all errors to Zerolog & sending the request's errors (not including panics) to the client.
// If unset, the status code of the request will be set to 500 and the Content-Type to test/plain.
func Recovery(c *gin.Context) {
	// PanicUserMessage defines the message that will be sent to the user when a panic occured.
	var PanicUserMessage = "server panicked"
	defer func() {
		if err := recover(); err != nil {
			if errCast, ok := err.(error); ok {
				log.Warn().Err(errCast).Str("path", c.FullPath()).Msg("request panicked")
			} else {
				log.Warn().Interface("error", err).Str("path", c.FullPath()).Msg("request panicked")
			}
			c.Error(errors.New(PanicUserMessage))
		}
		if len(c.Errors) > 0 {
			if !c.Writer.Written() {
				c.Status(http.StatusInternalServerError)
				c.Header("Content-Type", "text/plain")
			}
			if !c.Writer.Written() || c.Writer.Size() == 0 {
				errs := make([]error, len(c.Errors))
				for i, err := range c.Errors {
					errs[i] = err
					c.Writer.WriteString(err.Error() + "\n")
				}
				log.Debug().Errs("errors", errs).Str("path", c.FullPath()).Msg("request includes error(s)")
			}
		}
	}()
	c.Next()
}
