package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

// Default key names for the default fields
const (
	defaultTimestampFormat = time.RFC3339
	FieldKeyMsg            = "msg"
	FieldKeyLevel          = "level"
	FieldKeyTime           = "time"
)

// JSONFormatter formats logs into parsable json
type JSONLogger struct {
	// TimestampFormat sets the format used for marshaling timestamps.
	// The format to use is the same than for time.Format or time.Parse from the standard
	// library.
	// The standard Library already provides a set of predefined format.
	TimestampFormat   string
	// DisableTimestamp allows disabling automatic timestamps in output
	DisableTimestamp  bool
	// DisableHTMLEscape allows disabling html escaping in output
	DisableHTMLEscape bool
	// PrettyPrint will indent all json logs
	PrettyPrint       bool
}

// Format renders a single log entry
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
