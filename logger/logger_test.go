package logger

import (
	"testing"

	"github.com/coderc/onlinedisk-util/config"
)

func TestLogger(t *testing.T) {
	loggerConfig := config.Config{
		LoggerConfig: config.LoggerConfigStruct{
			Level: "debug",
			Path:  "./test_data/logs.log",
		},
	}
	GetLogger().Init(&loggerConfig.LoggerConfig)
	GetLogger().Debug("debug info")
	GetLogger().Info("info info")
}
