// Package log implements logger.
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

// Log is the global logger.
var Log *zerolog.Logger

// Logger returns the global logger.
func Logger() *zerolog.Logger {
	return Log
}

// init initializes the logger.
func init() {
	zerolog.SetGlobalLevel(setLogLevel())
	l := zerolog.New(os.Stdout).
		With().
		Timestamp().
		Caller().
		Logger()
	Log = &l
	Log.Info().Interface("log_level", zerolog.GlobalLevel()).Msg("logger configured")
}

// below are log levels list:
//  panic (zerolog.PanicLevel, 5)
//  fatal (zerolog.FatalLevel, 4)
//  error (zerolog.ErrorLevel, 3)
//  warn (zerolog.WarnLevel, 2)
//  info (zerolog.InfoLevel, 1)
//  debug (zerolog.DebugLevel, 0)
//  trace (zerolog.TraceLevel, -1)
// see https://github.com/rs/zerolog#leveled-logging
func setLogLevel() zerolog.Level {
	if ll, lookup := os.LookupEnv("EDT_LOG_LEVEL"); lookup {
		switch strings.ToLower(ll) {
		// case "debug":
		// 	return zerolog.DebugLevel
		case "disable":
			return zerolog.Disabled
		}
	}
	return zerolog.InfoLevel // assume everything except debug logs
}

// PrintfLogger holds logger.
type PrintfLogger struct {
	log *zerolog.Logger
}

// GetPrintfLogger returns the global logger.
func GetPrintfLogger(log *zerolog.Logger) *PrintfLogger {
	return &PrintfLogger{log: log}
}

// Printf implementation of the Logger interface.
func (l *PrintfLogger) Printf(format string, args ...interface{}) {
	l.log.Info().Msg(fmt.Sprintf(format, args...))
}
