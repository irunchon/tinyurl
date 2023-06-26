package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func Initialize(logLevel string) {
	var loggerConfig = zap.NewProductionConfig()

	parsedLogLevel, parsingError := zap.ParseAtomicLevel(logLevel)

	if parsingError == nil {
		loggerConfig.Level = parsedLogLevel
	}

	Logger = zap.Must(loggerConfig.Build())
	defer Logger.Sync()
}
