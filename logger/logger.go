package logger

import (
	"context"
	"flights-master/settings"
	"log/slog"
	"os"
)

type Logger struct {
	*slog.Logger
	level slog.Level
}

var defaultLogger *Logger

func init() {
}

func (l *Logger) WithError(err error) *Logger {
	n := l.With(slog.Attr{
		Key:   "error",
		Value: slog.AnyValue(err),
	})

	return &Logger{n, l.level}
}

func (l *Logger) WithAny(k string, val any) *Logger {
	n := l.With(slog.Any(k, val))
	return &Logger{n, l.level}
}

func (l *Logger) Panic(msg string) {
	l.Log(context.Background(), slog.Level(l.level), msg)
	panic(msg)
}

func New(s *settings.Settings) *Logger {
	level := slog.Level(s.LogLevel)

	opts := &slog.HandlerOptions{Level: level}
	base := slog.New(slog.NewJSONHandler(os.Stdout, opts))

	defaultLogger = &Logger{base, level}
	return defaultLogger
}

func FromContext(ctx context.Context) *Logger {
	return defaultLogger
}
