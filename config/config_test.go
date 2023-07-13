package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	configFilePath := "./test_data/config.yaml"
	err := GetConfig().Init(configFilePath)
	assert.NoError(t, err)
	assert.Equal(t, 8080, GetConfig().Port)
	assert.Equal(t, "debug", GetConfig().Level)
	assert.Equal(t, "./test_data/logs.log", GetConfig().Path)
}
