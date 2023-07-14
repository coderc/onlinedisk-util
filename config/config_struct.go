package config

type Config struct {
	ServerConfig ServerConfigStruct `yaml:"server"`
	LoggerConfig LoggerConfigStruct `yaml:"logger"`
	DBConfig     DBConfigStruct     `yaml:"db"`
}

// ServerConfigStruct server配置
type ServerConfigStruct struct {
	Port int `yaml:"port"`
}

// LoggerConfigStruct logger配置
type LoggerConfigStruct struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type DBConfigStruct struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DbName   string `yaml:"database"`

	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`

	ConnMaxLifetime int `yaml:"conn_max_lifetime"`
}

type RSAConfigStruct struct {
	PrivateKeyPath string `yaml:"private_key"`
	PublicKeyPath  string `yaml:"public_key"`
}
