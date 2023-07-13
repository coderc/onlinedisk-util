package logger

import (
	"testing"

	configStruct "github.com/coderc/onlinedisk-util/config"
)

func TestLogger(t *testing.T) {
	loggerConfig := configStruct.Config{
		LoggerConfig: configStruct.LoggerConfigStruct{
			Level: "debug",
			Path:  "./test_data/logs.log",
		},
	}
	GetLogger().Init(&loggerConfig.LoggerConfig)
	GetLogger().Debug("debug info")
	GetLogger().Info("info info")
}
