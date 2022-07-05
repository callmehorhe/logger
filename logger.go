package logger

import (
	"io"
	"time"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	entry *log.Entry
}

// Creates a new logger/
func New() Logger {
	l := log.New()
	return Logger{entry: log.NewEntry(l)}
}

// AddHook adds a hook to the logger hooks.
func (l *Logger) AddHook(hook log.Hook) {
	l.entry.Logger.AddHook(hook)
}

// SetFormatter sets the logger formatter.
func (l *Logger) SetFormatter(formatter log.Formatter) {
	l.entry.Logger.SetFormatter(formatter)
}

const (
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel log.Level = iota + 2
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

// SetLevel sets the logger level.
func (l *Logger) SetLevel(level log.Level) {
	l.entry.Logger.SetLevel(level)
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (log.Level, error) {
	return log.ParseLevel(lvl)
}

// Data type, used to pass to `WithFields`.
type Data map[string]interface{}

// WithField allocates a new entry and adds a field to it.
// If you want multiple fields, use `WithFields`.
func (l *Logger) WithField(key string, value interface{}) *Logger {
	return l.WithFields(Data{
		key: value,
	})
}

// WithFields add a map of fields to the Logger.
func (l *Logger) WithFields(fields Data) *Logger {
	data := make(map[string]interface{}, len(l.entry.Data)+len(fields))

	for k, v := range l.entry.Data {
		data[k] = v
	}
	for k, v := range fields {
		data[k] = v
	}

	return &Logger{
		entry: &log.Entry{
			Logger:  l.entry.Logger,
			Data:    data,
			Time:    l.entry.Time,
			Context: l.entry.Context,
		},
	}
}

// WithLatency adds latency to the logger.
//
// Use time.Duration's constant values to set format.
// For example:
// 	l.WithLatnecy(12345678, time.Hour)
func (l *Logger) WithLatency(latency time.Duration, format time.Duration) *Logger {
	var lat float64
	switch format {
	case time.Nanosecond:
		lat = float64(latency.Nanoseconds())
	case time.Microsecond:
		lat = float64(latency.Microseconds())
	case time.Millisecond:
		lat = float64(latency.Milliseconds())
	case time.Second:
		lat = latency.Seconds()
	case time.Minute:
		lat = latency.Minutes()
	case time.Hour:
		lat = latency.Hours()
	default:
		lat = float64(latency)
	}

	return l.WithField("latency", lat)
}

// SetOutput sets the standard logger output.
func (l *Logger) SetOutput(w io.Writer) {
	l.entry.Logger.SetOutput(w)
}

// Log logs a messege on the standard logger.
func (l *Logger) Log(level log.Level, args ...interface{}) {
	l.entry.Log(level, args)
}

// Debug logs a message at level Debug on the standard logger.
func (l *Logger) Debug(args ...interface{}) {
	l.entry.Log(log.DebugLevel, args...)
}

// Info logs a message at level Info on the standard logger.
func (l *Logger) Info(args ...interface{}) {
	l.entry.Log(log.InfoLevel, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l *Logger) Warn(args ...interface{}) {
	l.entry.Log(log.WarnLevel, args...)
}

// Error logs a message at level Error on the standard logger.
func (l *Logger) Error(args ...interface{}) {
	l.entry.Log(log.ErrorLevel, args...)
}

// Logf logs a messege on the standard logger.
func (l *Logger) Logf(level log.Level, format string, args ...interface{}) {
	l.entry.Logf(level, format, args)
}

// Debugf logs a message at level Debug on the standard logger.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.entry.Logf(log.DebugLevel, format, args...)
}

// Infof logs a message at level Info on the standard logger.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.entry.Logf(log.InfoLevel, format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.entry.Logf(log.WarnLevel, format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.entry.Logf(log.ErrorLevel, format, args...)
}
