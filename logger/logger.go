package logger

import (
	"fmt"
	"os"
	"sync"

	configStruct "github.com/coderc/onlinedisk-util/config_struct"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once     sync.Once
	instance *Logger
)

type Logger struct {
	zap *zap.Logger
}

func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{}
		instance.zap, _ = zap.NewProduction()
	})
	return instance
}

func (l *Logger) Init(config *configStruct.LoggerConfig) {
	file, err := os.Create(config.Path)
	if err != nil {
		msg := fmt.Sprintf(CreateLogFileError, err)
		panic(msg)
	}

	var level zapcore.Level
	switch config.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	case "panic":
		level = zap.PanicLevel
	default:
		level = zap.InfoLevel
	}

	// 设置 zapcore.WriteSyncer
	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // 编码器配置
		writeSyncer, // 输出到文件
		level,       // 日志级别
	)
	l.zap = zap.New(core, zap.AddCaller())
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.zap.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.zap.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.zap.Fatal(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...zap.Field) {
	l.zap.Panic(msg, fields...)
}

func (l *Logger) Sync() {
	l.zap.Sync()
}

const (
	CreateLogFileError = "create log file error: %v"
)
