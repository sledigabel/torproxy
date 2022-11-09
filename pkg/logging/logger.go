package logging

import (
	"go.uber.org/zap"
)

// var logger, _ = zap.NewProduction()
var logger *zap.Logger

func GetLogger() (*zap.Logger, error) {
	var err error
	if logger == nil {
		logger, err = zap.NewProduction()
	}
	return logger, err
}
