package logger

import (
	"io"
	"time"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	entry *log.Entry
}

func New() Logger {
	l := log.New()
	return Logger{entry: log.NewEntry(l)}
}

func (l *Logger) AddHook(hook log.Hook) {
	l.entry.Logger.AddHook(hook)
}

type Data map[string]interface{}

// Add a map of fields to the Logger.
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
		lat = float64(time.Duration(latency).Nanoseconds())
	case time.Microsecond:
		lat = float64(time.Duration(latency).Microseconds())
	case time.Millisecond:
		lat = float64(time.Duration(latency).Milliseconds())
	case time.Second:
		lat = time.Duration(latency).Seconds()
	case time.Minute:
		lat = time.Duration(latency).Minutes()
	case time.Hour:
		lat = time.Duration(latency).Hours()
	default:
		lat = float64(latency)
	}

	return l.WithFields(Data{
		"latency": lat,
	})
}

// SetOutput sets the standard logger output.
func (l *Logger) SetOutput(w io.Writer) {
	l.entry.Logger.SetOutput(w)
}

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
