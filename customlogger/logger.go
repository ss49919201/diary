package customlogger

import (
	"log/slog"
	"os"
	"sync"
)

type Logger interface {
	Info(string, ...any)
}

type logger struct {
	slogLogger *slog.Logger
}

var newLogger = sync.OnceValue(func() Logger {
	l := logger{}
	l.slogLogger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	return l
})

func NewLogger() Logger {
	return newLogger()
}

func (l logger) Info(msg string, attr ...any) {
	l.slogLogger.Info(msg, attr...)
}
