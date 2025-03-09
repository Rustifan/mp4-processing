package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	log *zap.SugaredLogger
}

func newZapLogger(serviceName string) (*zapLogger, error) {
	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err := config.Build(zap.AddCallerSkip(1), zap.Fields(
		zap.String("service", serviceName),
	))
	if err != nil {
		return nil, err
	}
	return &zapLogger{
		log: logger.Sugar(),
	}, nil
}

func (logger *zapLogger) Debug(msg string, keyvals ...interface{}) {
	logger.log.Debugw(msg, keyvals...)
}

func (logger *zapLogger) Info(msg string, keyvals ...interface{}) {
	logger.log.Infow(msg, keyvals...)
}

func (logger *zapLogger) Error(msg string, keyvals ...interface{}) {
	logger.log.Errorw(msg, keyvals...)
}

func (logger *zapLogger) Fatal(msg string, keyvals ...interface{}) {
	logger.log.Fatalw(msg, keyvals...)
}
