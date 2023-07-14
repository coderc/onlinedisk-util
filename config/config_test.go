package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configFilePath := "./test_data/config.yaml"
	GetConfig().Init(configFilePath)
	assert.Equal(t, 8080, GetConfig().ServerConfig.Port)
	assert.Equal(t, "debug", GetConfig().LoggerConfig.Level)
	assert.Equal(t, "./test_data/logs.log", GetConfig().LoggerConfig.Path)
}
