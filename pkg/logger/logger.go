package logger

import (
	"log/slog"
	"sync"
)

type LInstance struct {
	logger *slog.Logger
}

var (
	loggerInstance     *LInstance
	loggerInstanceOnce sync.Once
)

func Instance() *LInstance {
	if loggerInstance != nil {
		return loggerInstance
	}
	// TODO: add logging level
	loggerInstanceOnce.Do(func() { loggerInstance = &LInstance{logger: slog.Default()} })
	return loggerInstance
}

func (l *LInstance) Info(msg *Message) {
	msgBytes, err := msg.ToBytes()
	if err != nil {
		l.logger.Error("failed to log with error: " + err.Error())
		return
	}
	l.logger.Info(string(msgBytes))
}

func (l *LInstance) Error(msg *Message) {
	msgBytes, err := msg.ToBytes()
	if err != nil {
		l.logger.Error("failed to log with error: " + err.Error())
		return
	}
	l.logger.Error(string(msgBytes))
}

func (l *LInstance) Warn(msg *Message) {
	msgBytes, err := msg.ToBytes()
	if err != nil {
		l.logger.Error("failed to log with error: " + err.Error())
		return
	}
	l.logger.Warn(string(msgBytes))
}
