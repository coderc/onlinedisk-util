package config

type Config struct {
	Version           string                  `yaml:"version"`
	ServerConfig      ServerConfigStruct      `yaml:"server"`
	LoggerConfig      LoggerConfigStruct      `yaml:"logger"`
	DBConfig          DBConfigStruct          `yaml:"db"`
	RedisConfig       RedisConfigStruct       `yaml:"redis"`
	UserServiceConfig UserServiceConfigStruct `yaml:"user_service"`
	RabbitMQConfig    RabbitMQConfigStruct    `yaml:"rabbitmq"`
	OssConfig         OssConfigStruct         `yaml:"oss"`
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

type RedisConfigStruct struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`

	MaxIdle     int `yaml:"max_idle"`
	MaxActive   int `yaml:"max_active"`
	IdleTimeout int `yaml:"idle_timeout"`
}

type UserServiceConfigStruct struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type RabbitMQConfigStruct struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type OssConfigStruct struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"access_key_id"`
	AccessKeySecret string `yaml:"access_key_secret"`
	BucketName      string `yaml:"bucket_name"`
}
