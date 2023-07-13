package config

// ServerConfigStruct server配置
type ServerConfigStruct struct {
	Port int `yaml:"port"`
}

// LoggerConfigStruct logger配置
type LoggerConfigStruct struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}
