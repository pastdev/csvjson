package c2j

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ Writer = &JSONObjectWriter{}

type JSONObjectWriter struct {
	w       io.Writer
	lastRow map[string]string
}

func (w JSONObjectWriter) Begin() error {
	_, err := w.w.Write([]byte("["))
	if err != nil {
		return fmt.Errorf("begin: %w", err)
	}
	return nil
}

func (w JSONObjectWriter) End() error {
	_, err := w.w.Write([]byte("]"))
	if err != nil {
		return fmt.Errorf("end: %w", err)
	}
	return nil
}

func (w *JSONObjectWriter) WriteRecord(r *Record) error {
	if len(w.lastRow) == 0 {
		w.lastRow = make(map[string]string, len(r.Header))
	} else {
		_, err := w.w.Write([]byte(","))
		if err != nil {
			return fmt.Errorf("write record separator: %w", err)
		}
	}

	for i := range r.Header {
		w.lastRow[r.Header[i]] = r.Row[i]
	}

	data, err := json.Marshal(w.lastRow)
	if err != nil {
		return fmt.Errorf("marshal record: %w", err)
	}

	_, err = w.w.Write(data)
	if err != nil {
		return fmt.Errorf("write record: %w", err)
	}

	return nil
}

func NewJSONObjectWriter(w io.Writer) *JSONObjectWriter {
	return &JSONObjectWriter{w: w}
}
