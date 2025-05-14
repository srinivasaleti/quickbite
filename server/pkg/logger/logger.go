package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Info(msg string, keysAndValues ...any)
	Error(err error, msg string, keysAndValues ...any)
}

func createZapLogger(level zap.AtomicLevel) *zap.Logger {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.Config{
		Level:             level,
		Development:       false,
		DisableCaller:     false,
		DisableStacktrace: false,
		Encoding:          "json",
		OutputPaths: []string{
			"stdout",
		},
		ErrorOutputPaths: []string{
			"stdout",
		},
		EncoderConfig: encoderCfg,
	}

	return zap.Must(config.Build())
}

// Logger wraps zap.Logger for structured logging with Info, Error, and Debug methods.
type Logger struct {
	ZapLogger *zap.Logger
}

func (l *Logger) Info(msg string, keysAndValues ...any) {
	if l.ZapLogger != nil {
		fields := l.transformToZapFields(keysAndValues...)
		l.ZapLogger.Info(msg, fields...)
	}
}

func (l *Logger) Error(err error, msg string, keysAndValues ...any) {
	if l.ZapLogger != nil {
		fields := l.transformToZapFields(keysAndValues...)
		fields = append(fields, zap.Error(err))
		l.ZapLogger.Error(msg, fields...)
	}
}

// transformToZapFields converts key-value pairs to a slice of zap.Field.
// It assumes that the keys are always strings and values can be of any type.
func (l *Logger) transformToZapFields(keysAndValues ...any) []zap.Field {
	var fields []zap.Field
	for i := 0; i < len(keysAndValues)-1; i += 2 {
		if key, ok := keysAndValues[i].(string); ok {
			fields = append(fields, zap.Any(key, keysAndValues[i+1]))
		}
	}
	return fields
}

func NewLogger(level string) (ILogger, error) {
	atomicLevel, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return nil, err
	}
	return &Logger{ZapLogger: createZapLogger(atomicLevel)}, nil
}
