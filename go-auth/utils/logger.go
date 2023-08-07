package utils

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerEnvironment string

const (
	DevelopmentLoggerEnvironment LoggerEnvironment = "development"
	ProductionLoggerEnvironment  LoggerEnvironment = "production"
)

func GetLogger(environment_provided string) *zap.SugaredLogger {
	// Select logger environment
	var environment LoggerEnvironment

	if strings.ToLower(environment_provided) == string(DevelopmentLoggerEnvironment) {
		environment = DevelopmentLoggerEnvironment
	} else {
		environment = ProductionLoggerEnvironment
	}

	// Initialize logger
	var logger_base *zap.Logger

	if environment == DevelopmentLoggerEnvironment {
		config := zap.NewDevelopmentConfig()
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		logger, _ := config.Build()
		logger_base = logger
	} else if environment == ProductionLoggerEnvironment {
		config := zap.NewProductionConfig()

		logger, _ := config.Build()
		logger_base = logger
	}

	return logger_base.Sugar()
}
