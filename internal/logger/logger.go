package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"mfo-service/internal/config"
)

type Logger = *zap.SugaredLogger

func New(cfg *config.LoggerConfig) (Logger, error) {
	l := zap.New(
		zapcore.NewCore(
			zapcore.NewJSONEncoder(zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				CallerKey:      "caller",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
			zapcore.AddSync(os.Stdout),
			zap.NewAtomicLevelAt(cfg.GetLevel()),
		),
		zap.AddCaller(),
	)

	zap.ReplaceGlobals(l)
	return l.Sugar(), nil
}
