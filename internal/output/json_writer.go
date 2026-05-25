package output

import (
	"encoding/json"
	"io"

	"github.com/drone-relay/drone-relay-agent/internal/model"
)

type JSONWriter struct {
	writer io.Writer
}

func NewJSONWriter(writer io.Writer) *JSONWriter {
	return &JSONWriter{
		writer: writer,
	}
}

func (w *JSONWriter) Write(metrics model.HostMetrics) error {
	encoder := json.NewEncoder(w.writer)
	return encoder.Encode(metrics)
}

func (w *JSONWriter) WriteBatch(metrics []model.HostMetrics) error {
	encoder := json.NewEncoder(w.writer)
	return encoder.Encode(metrics)
}