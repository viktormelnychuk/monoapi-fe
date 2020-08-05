package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func New() *zap.SugaredLogger {
	zapConfig := zap.NewDevelopmentConfig()

	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zapLogger, err := zapConfig.Build()
	if err != nil {
		panic(fmt.Errorf("init loger %v", err))
	}
	return zap.New(zapLogger.Core()).Sugar()
}
