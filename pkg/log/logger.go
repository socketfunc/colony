package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type options struct {
	level string
}

type Option func(*options)

func NewLogger(opts ...Option) (*zap.Logger, error) {
	o := &options{level: "info"}
	for i := range opts {
		opts[i](o)
	}
	level := new(zapcore.Level)
	if err := level.Set(o.level); err != nil {
		return nil, err
	}
	return newJSONConfig(*level).Build()
}

func newJSONConfig(level zapcore.Level) *zap.Config {
	return &zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Encoding:    "json",
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		EncoderConfig:    newEncodeConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func newEncodeConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochNanosTimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
