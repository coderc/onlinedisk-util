package configstruct

// ServerConfig server配置
type ServerConfig struct {
	Port int `yaml:"port"`
}

// LoggerConfig logger配置
type LoggerConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}
