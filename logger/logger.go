package logger

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Zaplog struct {
	logger *zap.Logger
}

var Zaplogger Zaplog

func InitLogger() {
	var err error
	var cfg zap.Config = zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	cfg.EncoderConfig.FunctionKey = "func"
	Zaplogger.logger, err = cfg.Build()
	//SugarLog = Zaplogger.Sugar()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() {
		err := Zaplogger.logger.Sync()
		if err != nil {
			fmt.Printf("can't sync zap logger: %v", err)
		}
	}()
	//defer Zaplogger.logger.Sync()
}

func (z *Zaplog) Printf(message string, fields ...interface{}) {
	z.logger.Info(message)
}

func Info(message string, fields ...zap.Field) {
	Zaplogger.logger.Info(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Zaplogger.logger.Warn(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Zaplogger.logger.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Zaplogger.logger.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	Zaplogger.logger.Fatal(message, fields...)
}
