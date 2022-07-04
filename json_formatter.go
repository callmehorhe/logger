package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = "msg"
	FieldKeyLevel          = "level"
	FieldKeyTime           = "time"
)

type JSONLogger struct {
	TimestampFormat   string
	DisableTimestamp  bool
	DisableHTMLEscape bool
	PrettyPrint       bool
}

func (j *JSONLogger) Format(entry *log.Entry) ([]byte, error) {
	data := make(Data, len(entry.Data)+3)

	for k, v := range entry.Data {
		data[k] = v
	}
	timestampFormat := j.TimestampFormat
	if timestampFormat == "" {
		timestampFormat = defaultTimestampFormat
	}
	if !j.DisableTimestamp {
		data[FieldKeyTime] = entry.Time.Format(time.RFC3339)
	}
	data[FieldKeyMsg] = entry.Message
	data[FieldKeyLevel] = entry.Level.String()

	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	encoder := json.NewEncoder(b)
	encoder.SetEscapeHTML(!j.DisableHTMLEscape)
	if j.PrettyPrint {
		encoder.SetIndent("", "  ")
	}
	if err := encoder.Encode(data); err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %w", err)
	}

	return b.Bytes(), nil
}
