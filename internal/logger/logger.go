package logger

import (
	"time"

	"go.uber.org/zap"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	return &Logger{
		logger: logger,
	}
}

func (l *Logger) Info(msg, info string) {
	l.logger.Info(msg,
		zap.String("time", getTime()),
		zap.String("info", info),
	)
}

func (l *Logger) Fatal(msg string, err error) {
	l.logger.Fatal(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Error(msg string, err error) {
	l.logger.Error(msg,
		zap.String("time", getTime()),
		zap.String("error", err.Error()),
	)
}

func (l *Logger) Sync() {
	l.logger.Sync()
}

func getTime() string {
	return time.Now().Format(time.RFC3339Nano)
}
