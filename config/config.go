package config

import (
	"sync"

	"github.com/coderc/onlinedisk-util/utils"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerConfig ServerConfigStruct `yaml:"server"`
	LoggerConfig LoggerConfigStruct `yaml:"logger"`
}

var (
	once sync.Once
	cfg  *Config
)

// GetConfig 获取配置
func GetConfig() *Config {
	once.Do(func() {
		cfg = &Config{}
	})
	return cfg
}

// Init 初始化配置
func (cfg *Config) Init(configFilePath string) error {
	bytes, err := utils.GetFileBytes(configFilePath)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		return err
	}

	return nil
}
