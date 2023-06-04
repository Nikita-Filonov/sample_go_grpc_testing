package logger

import (
	"sample_go_grpc_testing/utils/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var MapConfLevelZapLevel = map[config.LogLevel]zapcore.Level{
	config.DebugLogLevel:   zap.DebugLevel,
	config.InfoLogLevel:    zap.InfoLevel,
	config.WarningLogLevel: zap.WarnLevel,
	config.ErrorLogLevel:   zap.ErrorLevel,
}

type Service interface {
	NewPrefix(prefix string) *CtxLogger
}

type loggerService struct {
	*zap.Logger
}

func (ls *loggerService) NewPrefix(prefix string) *CtxLogger {
	return &CtxLogger{ls.Logger.Named(prefix)}
}

func NewLoggerService(conf config.Config) (Service, error) {
	cfg := newConfig(conf.Logger)

	zapLogger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &loggerService{zapLogger}, err
}

func utcTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z0700"))
}

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger_name",
		CallerKey:      "caller_file",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     utcTimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func newConfig(conf config.Logger) zap.Config {
	cfg := zap.Config{
		Level:             zap.NewAtomicLevelAt(MapConfLevelZapLevel[conf.Level]),
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newEncoderConfig(),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
	}

	if conf.IsDevMode {
		cfg.Development = true
		cfg.Encoding = "console"
	}

	return cfg
}
