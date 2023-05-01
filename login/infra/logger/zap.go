package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func NewLogger() *zap.Logger {
	config := SetConfig()

	enc := zapcore.NewJSONEncoder(config.EncoderConfig)
	sink := zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   "/go/src/login/log/develop.log",
			MaxSize:    5,
			MaxBackups: 3,
			MaxAge:     1,
		},
	)

	core := zapcore.NewTee(
		zapcore.NewCore(enc, sink, config.Level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(config.EncoderConfig), zapcore.AddSync(os.Stdout), config.Level),
	)
	logger := zap.New(core)

	return logger
}

func SetConfig() zap.Config {
	config := zap.Config{
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Development: true,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     JSTTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	return config
}

func JSTTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	const layout = "2006-01-02T15:04:05"

	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	enc.AppendString(t.In(jst).Format(layout))
}
