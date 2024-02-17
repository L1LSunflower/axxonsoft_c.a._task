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
	if loggerInstance == nil {
		loggerInstanceOnce.Do(func() { loggerInstance = &LInstance{logger: slog.Default()} })
	}
	return loggerInstance
}

func (l *LInstance) Info(msg *Message) {
	l.logger.Info(msg.Message, msg.FullMessage, msg.Datetime, msg.ExtraData)
}

func (l *LInstance) Error(msg *Message) {
	l.logger.Info(msg.Message, msg.FullMessage, msg.Datetime, msg.ExtraData)
}

func (l *LInstance) Warn(msg *Message) {
	l.logger.Info(msg.Message, msg.FullMessage, msg.Datetime, msg.ExtraData)
}
