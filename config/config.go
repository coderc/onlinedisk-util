package config

import (
	"sync"

	"github.com/coderc/onlinedisk-util/utils"
	"gopkg.in/yaml.v3"
)


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
func (cfg *Config) Init(configFilePath string) {
	bytes, err := utils.GetFileBytes(configFilePath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(bytes, cfg)
	if err != nil {
		panic(err)
	}
}
