package log

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	log, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return log
}
