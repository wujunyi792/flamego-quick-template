package logx

import (
	"github.com/wujunyi792/flamego-quick-template/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

// NameSpace - 提供带有模块命名空间的logger
func NameSpace(name string) *zap.SugaredLogger {
	return zap.S().Named(name)
}

func getLogWriter() zapcore.WriteSyncer {
	if config.GetConfig().LogPath == "" {
		config.GetConfig().LogPath = "app.log"
		print("LogPath 未设置, 使用默认值app.log")
	}
	lj := &lumberjack.Logger{
		Filename:   config.GetConfig().LogPath,
		MaxSize:    5,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
	return zapcore.AddSync(lj)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func Init(level zapcore.LevelEnabler) {
	writeSyncer := getLogWriter()
	if level == zapcore.DebugLevel {
		writeSyncer = zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout))
	}
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level)
	zap.ReplaceGlobals(zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel)))
}
