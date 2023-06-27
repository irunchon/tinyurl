package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func Initialize(logLevel string) {
	var loggerConfig = zap.NewProductionConfig()

	parsedLogLevel, parsingError := zap.ParseAtomicLevel(logLevel)
	if parsingError == nil {
		loggerConfig.Level = parsedLogLevel
	}

	loggerConfig.DisableCaller = true
	Logger = zap.Must(loggerConfig.Build())
	defer Logger.Sync()

	if parsingError != nil {
		Logger.Error(
			fmt.Sprintf("failed to parse log level: %v (log level set to info)",
				parsingError))
	}
}
