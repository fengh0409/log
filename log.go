package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Sugaredlogger is alias of zap.Sugaredlogger
type SugaredLogger = zap.SugaredLogger

// Logger is alias of zap.Logger
type Logger = zap.Logger

// Field is alias of zap.Field
type Field = zap.Field

// Level is alias of zapcore.Level
type Level = zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in production.
	DebugLevel Level = zap.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel Level = zap.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual human review.
	WarnLevel Level = zap.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel Level = zap.ErrorLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel
)

const (
	// ConsoleEncoder output the log with console
	ConsoleEncoder = "console"
	// JSONEncoder output the log with json format
	JSONEncoder = "json"
)

// _logger is a global Log
var _logger = New().Build()

// New initialize the options for building
func New(options ...Option) *Options {
	o := &Options{}
	for _, option := range options {
		option(o)
	}

	if o.writer == nil {
		o.writer = os.Stderr
	}
	if o.level == 0 {
		o.level = InfoLevel
	}
	if o.encoding == "" {
		o.encoding = ConsoleEncoder
	}

	return o
}

// Build construct a Log from Options
func (o *Options) Build() *Log {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	if o.encoding == JSONEncoder {
		encoderConfig.EncodeTime = zapcore.EpochTimeEncoder
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	zapCore := zapcore.NewCore(
		encoder,
		zapcore.AddSync(o.writer),
		zapcore.Level(o.level),
	)

	return &Log{
		logger: zap.New(zapCore, zap.AddCaller(), zap.AddCallerSkip(1)),
	}
}

// SetOptions set options
func SetOptions(options ...Option) {
	_logger = New(options...).Build()
}

// Debug logs a message at level Debug on the standard logger
func Debug(args ...interface{}) {
	_logger.Sugar().Debug(args...)
}

// Info logs a message at level Info on the standard logger
func Info(args ...interface{}) {
	_logger.Sugar().Info(args...)
}

// Warn logs a message at level Warn on the standard loggero
func Warn(args ...interface{}) {
	_logger.Sugar().Warn(args...)
}

// Error logs a message at level Error on the standard loggero
func Error(args ...interface{}) {
	_logger.Sugar().Error(args...)
}

// Fatal logs a message at level Fatal on the standard loggero
func Fatal(args ...interface{}) {
	_logger.Sugar().Fatal(args...)
}

// Debugf logs a message at level Debug on the standard loggero
func Debugf(template string, args ...interface{}) {
	_logger.Sugar().Debugf(template, args...)
}

// Infof logs a message at level Info on the standard loggero
func Infof(template string, args ...interface{}) {
	_logger.Sugar().Infof(template, args...)
}

// Warnf logs a message at level Warn on the standard loggero
func Warnf(template string, args ...interface{}) {
	_logger.Sugar().Warnf(template, args...)
}

// Errorf logs a message at level Error on the standard loggero
func Errorf(template string, args ...interface{}) {
	_logger.Sugar().Errorf(template, args...)
}

// Fatalf logs a message at level Fatal on the standard loggero
func Fatalf(template string, args ...interface{}) {
	_logger.Sugar().Fatalf(template, args...)
}

// Sync calls the sugar's Sync method
func Sync() error {
	return _logger.Sugar().Sync()
}

// Log wraps zap logger
type Log struct {
	logger *Logger
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Log) Debug(msg string, fields ...Field) {
	l.logger.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Log) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Log) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Log) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Log) Fatal(msg string, fields ...Field) {
	l.logger.Fatal(msg, fields...)
}

// Sync calls the logger's Sync method
func (l *Log) Sync() error {
	return l.logger.Sync()
}

// Sugar calls the logger's Sugar method
func (l *Log) Sugar() *SugaredLogger {
	return l.logger.Sugar()
}

var (
	// Any is alias of zap.Any Field
	Any = zap.Any
	// Bool is alias of zap.Bool Field
	Bool = zap.Bool
	// String is alias of zap.String Field
	String = zap.String
	// Float32 alias of zap.Float32 Field
	Float32 = zap.Float32
	// Float64 is alias of zap.Float64 Field
	Float64 = zap.Float64
	// Int is alias of zap.Int Field
	Int = zap.Int
	// Int8 is alias of zap.Int8 Field
	Int8 = zap.Int8
	// Int16 is alias of zap.Intl6 Field
	Int16 = zap.Int16
	// Int32 is alias of zap.Int32 Field
	Int32 = zap.Int32
	// Int64 is alias of zap.Int64 Field
	Int64 = zap.Int64
	// Uint is alias of zap.Uint Field
	Uint = zap.Uint
	// Uint8 is alias of zap.Uint8 Field
	Uint8 = zap.Uint8
	// Vint16 is alias of zap.Uintl6 Field
	Uint16 = zap.Uint16
	// Uint32 is alias of zap.Uint32 Field
	Uint32 = zap.Uint32
	// Uint64 is alias of zap.Uint64 Field
	Uint64 = zap.Uint64
	// Namespace is alias of zap.Namespace Field
	Namespace = zap.Namespace
)
