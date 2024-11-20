package logging

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	loggerContextKey contextKey = "logger"
)

const (
	Debug Level = 4 * (iota - 1)
	Info
	Warn
	Error
)

type contextKey string

// Level represents log severity.
type Level int

// Field represents Logger key-value field.
type Field struct {
	Key   string
	Value any
}

type Logger interface {
	Log(severity Level, msg string, args ...Field)
	Info(msg string, args ...Field)
	Error(msg string, args ...Field)
	Warn(msg string, args ...Field)
	Debug(msg string, args ...Field)
}

// DefaultLogger represents default logger.
type DefaultLogger struct {
	logger *zap.SugaredLogger
}

// NewDefaultLogger creates new instance of default logger.
func NewDefaultLogger() (DefaultLogger, error) {
	// Define the zap configuration
	loggerCfg := &zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	plain, err := loggerCfg.Build(zap.AddStacktrace(zap.DPanicLevel))
	if err != nil {
		return DefaultLogger{logger: zap.NewNop().Sugar()}, err
	}

	return DefaultLogger{logger: plain.Sugar()}, nil
}

// Log creates new log of provided severity with message and args.
func (d DefaultLogger) Log(level Level, msg string, args ...Field) {
	fields := make([]interface{}, len(args)*2)
	for idx, val := range args {
		fields[idx*2] = val.Key
		fields[idx*2+1] = val.Value
	}

	switch level {
	case Info:
		d.logger.Infow(msg, fields...)
	case Error:
		d.logger.Errorw(msg, fields...)
	case Warn:
		d.logger.Warnw(msg, fields...)
	case Debug:
		d.logger.Debugw(msg, fields...)
	}
}

// Info prints log of Info level.
func (d DefaultLogger) Info(msg string, args ...Field) {
	d.Log(Info, msg, args...)
}

// Error prints log of Error level.
func (d DefaultLogger) Error(msg string, args ...Field) {
	d.Log(Error, msg, args...)
}

// Warn prints log of Warn level.
func (d DefaultLogger) Warn(msg string, args ...Field) {
	d.Log(Warn, msg, args...)
}

// Debug prints log of Debug level.
func (d DefaultLogger) Debug(msg string, args ...Field) {
	d.Log(Debug, msg, args...)
}

// Any creates an instance of Field of any type value.
func Any(key string, value any) Field {
	return Field{Key: key, Value: value}
}

// WrapToContext injects logger into context to be passed within app layers.
func WrapToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerContextKey, logger)
}

// FromContext extracts Logger instance from Context or returns a default one.
func FromContext(ctx context.Context) Logger {
	value := ctx.Value(loggerContextKey)
	if value == nil {
		logger, err := NewDefaultLogger()
		if err != nil {
			panic("failed to create default logger during extracting from context")
		}

		return logger
	}

	logger, ok := value.(Logger)
	if !ok {
		logger, err := NewDefaultLogger()
		if err != nil {
			panic("failed to create default logger during extracting from context")
		}

		return logger
	}

	return logger
}

// Encoder config used in logger configuration.
var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "time",
	LevelKey:       "severity",
	NameKey:        "logger",
	CallerKey:      "caller",
	MessageKey:     "message",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    encodeLevel(),
	EncodeTime:     zapcore.RFC3339TimeEncoder,
	EncodeDuration: zapcore.MillisDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

// encodeLevel translates zapcore levels to specified provider cloud logging severities.
func encodeLevel() zapcore.LevelEncoder {
	return func(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		switch l {
		case zapcore.DebugLevel:
			enc.AppendString("DEBUG")
		case zapcore.InfoLevel:
			enc.AppendString("INFO")
		case zapcore.WarnLevel:
			enc.AppendString("WARNING")
		case zapcore.ErrorLevel:
			enc.AppendString("ERROR")
		case zapcore.DPanicLevel:
			enc.AppendString("CRITICAL")
		case zapcore.PanicLevel:
			enc.AppendString("ALERT")
		case zapcore.FatalLevel:
			enc.AppendString("EMERGENCY")
		}
	}
}
