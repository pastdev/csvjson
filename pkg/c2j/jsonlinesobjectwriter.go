package c2j

import (
	"encoding/json"
	"fmt"
	"io"
)

var _ Writer = &JSONLinesObjectWriter{}

type JSONLinesObjectWriter struct {
	w       io.Writer
	lastRow map[string]string
}

func (w JSONLinesObjectWriter) Begin() error {
	return nil
}

func (w JSONLinesObjectWriter) End() error {
	return nil
}

func (w *JSONLinesObjectWriter) WriteRecord(r *Record) error {
	if len(w.lastRow) == 0 {
		w.lastRow = make(map[string]string, len(r.Header))
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
	_, err = w.w.Write([]byte("\n"))
	if err != nil {
		return fmt.Errorf("write record separator: %w", err)
	}

	return nil
}

func NewJSONLinesObjectWriter(w io.Writer) *JSONLinesObjectWriter {
	return &JSONLinesObjectWriter{w: w}
}
