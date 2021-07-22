package libs

import (
	zap "go.uber.org/zap"
)

type _Logger struct {
	PAPLogger
}

type PAPLogger interface {
	Print() *zap.SugaredLogger
}

func (l *_Logger) Print() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()
	return sugar
}

var Logger *_Logger
