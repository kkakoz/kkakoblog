package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger(viper *viper.Viper) {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	var core zapcore.Core
	if viper.GetString("mode") == gin.DebugMode {
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel),
			zapcore.NewCore(consoleEncoder, writeSyncer, zapcore.DebugLevel),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	}
	//core = zapcore.RegisterHooks(core, func(entry zapcore.Entry) error {
	//	fmt.Println("entry = ", entry)
	//	return nil
	//})
	logger = zap.New(core)
	sugarLogger = logger.Sugar()
}

func L() *zap.Logger {
	return logger
}

func SL() *zap.SugaredLogger {
	return sugarLogger
}

func getEncoder() zapcore.Encoder {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(config)
}

func getLogWriter() zapcore.WriteSyncer {
	file, err := os.Create("./storage/log/one.log")
	if err != nil {
		fmt.Println("create file err")
	}
	return zapcore.AddSync(file)
}