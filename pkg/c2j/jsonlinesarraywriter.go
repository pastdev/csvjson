package c2j

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ Writer = &JSONLinesArrayWriter{}

type JSONLinesArrayWriter struct {
	w io.Writer
}

func (w JSONLinesArrayWriter) Begin() error {
	return nil
}

func (w JSONLinesArrayWriter) End() error {
	return nil
}

func (w *JSONLinesArrayWriter) WriteRecord(r *Record) error {
	data, err := json.Marshal(r.Row)
	if err != nil {
		return fmt.Errorf("marshal record: %w", err)
	}

	_, err = w.w.Write(data)
	if err != nil {
		return fmt.Errorf("write record: %w", err)
	}
	_, err = w.w.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("write record separator: %w", err)
	}

	return nil
}

func NewJSONLinesArrayWriter(w io.Writer) *JSONLinesArrayWriter {
	return &JSONLinesArrayWriter{w: w}
}
