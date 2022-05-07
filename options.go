package log

import (
	"io"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

var _levelConvert = map[string]Level{
	"debug": DebugLevel,
	"info":  InfoLevel,
	"warn":  WarnLevel,
	"error": ErrorLevel,
	"fatal": FatalLevel,
}

// Options define some options for building a Log
type Options struct {
	writer   io.Writer
	level    Level
	encoding string
}

// Option define a func wraps Options
type Option func(*Options)

// WithWriter specify the writer where log to write
func WithWriter(writer io.Writer) Option {
	return func(o *Options) {
		o.writer = writer
	}
}

// withLevel specify the log level
func WithLevel(level Level) Option {
	return func(o *Options) {
		o.level = level
	}
}

// withLevelString convert the level to Level type, if not exists, we set it as InfoLevel default
func WithLevelString(levelString string) Option {
	level, ok := _levelConvert[strings.ToLower(levelString)]
	if !ok {
		level = InfoLevel
	}

	return WithLevel(level)
}

// WithEncoding specify the format of output, "json" and "console" are supported
func WithEncoding(encoding string) Option {
	return func(o *Options) {
		o.encoding = encoding
	}
}

// WithFileWriter support write log to local file, with the FileOption configures
func WithFileWriter(options ...FileOption) Option {
	fo := &FileOptions{}
	for _, option := range options {
		option(fo)
	}

	if fo.filename == "" {
		fo.filename = "/tmp/log/lumberjack.log"
	}
	if fo.maxSize == 0 {
		fo.maxSize = 200
	}
	if fo.maxAge == 0 {
		fo.maxAge = 7
	}
	if fo.maxBackups == 0 {
		fo.maxBackups = 10
	}
	compress := true
	if fo.compress == nil {
		fo.compress = &compress
	}

	writer := &lumberjack.Logger{
		Filename:   fo.filename,
		MaxSize:    fo.maxSize,
		MaxAge:     fo.maxAge,
		MaxBackups: fo.maxBackups,
		Compress:   *fo.compress,
		LocalTime:  true,
	}
	return WithWriter(writer)
}

// FileOptions define the file writer options
type FileOptions struct {
	// filename is the file to write logs to
	filename string
	// maxSize is the maximum size in MB of the file before it gets rotated
	maxSize int
	// maxAge is the maximum number days to retain old files based on the timestamp encoded in their filename
	maxAge int
	// maxBackups is the maximum number of old log files to retain
	maxBackups int
	// compress determines if the rotated files should be compressed using gzip
	compress *bool
}

// FileOption define a func wraps FileOptions
type FileOption func(*FileOptions)

// WithFilename specify the filename where log to write
func WithFilename(filename string) FileOption {
	return func(o *FileOptions) {
		o.filename = filename
	}
}

// WithMaxSize specify maximum size in MB of the file before it gets rotated
func WithMaxSize(maxSize int) FileOption {
	return func(o *FileOptions) {
		o.maxSize = maxSize
	}
}

// WithMaxAge specify the maximum number days to retain old files based on the timestamp encoded in their filename
func WithMaxAge(maxAge int) FileOption {
	return func(o *FileOptions) {
		o.maxAge = maxAge
	}
}

// WithMaxBackups specify the maximum number of old log files to retain
func WithMaxBackups(maxBackups int) FileOption {
	return func(o *FileOptions) {
		o.maxBackups = maxBackups
	}
}

// WithCompress specify if the rotated files should be compressed using gzip
func WithCompress(compress bool) FileOption {
	return func(o *FileOptions) {
		o.compress = &compress
	}
}
