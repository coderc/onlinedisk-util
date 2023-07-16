package logger

import (
	"fmt"
	"os"
	"sync"

	"github.com/coderc/onlinedisk-util/config"
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

func Zap() *zap.Logger {
	return GetLogger().zap
}

func (l *Logger) Init(config *config.LoggerConfigStruct) {
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

const (
	CreateLogFileError = "create log file error: %v"
)
