package c2j

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ Writer = &JSONArrayWriter{}

type JSONArrayWriter struct {
	w        io.Writer
	notFirst bool
}

func (w JSONArrayWriter) Begin() error {
	_, err := w.w.Write([]byte("["))
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}
	return nil
}

func (w JSONArrayWriter) End() error {
	_, err := w.w.Write([]byte("]"))
	if err != nil {
		return fmt.Errorf("end: %w", err)
	}
	return nil
}

func (w *JSONArrayWriter) WriteRecord(r *Record) error {
	if w.notFirst {
		_, err := w.w.Write([]byte(","))
		if err != nil {
			return fmt.Errorf("write record separator: %w", err)
		}
	} else {
		w.notFirst = true
	}

	data, err := json.Marshal(r.Row)
	if err != nil {
		return fmt.Errorf("marshal record: %w", err)
	}

	_, err = w.w.Write(data)
	if err != nil {
		return fmt.Errorf("write record: %w", err)
	}

	return nil
}

func NewJSONArrayWriter(w io.Writer) *JSONArrayWriter {
	return &JSONArrayWriter{w: w}
}
