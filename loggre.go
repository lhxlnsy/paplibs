package libs

import (
	zap "go.uber.org/zap"
)

type Logger struct {
	PAPLogger
}

type PAPLogger interface {
	Print() *zap.SugaredLogger
}

func (l *Logger) Print() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

var _Logger *Logger
