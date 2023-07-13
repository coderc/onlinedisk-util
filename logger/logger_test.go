package logger

import (
	"testing"

	configStruct "github.com/coderc/onlinedisk-util/config_struct"
)

func TestLogger(t *testing.T) {
	loggerConfig := configStruct.LoggerConfig{
		Level: "info",
		Path:  "./test_data/test.log",
	}
	GetLogger().Init(&loggerConfig)
	GetLogger().Debug("debug info")
	GetLogger().Info("info info")
}
