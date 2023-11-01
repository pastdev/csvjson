package c2j

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"strconv"
)

type csvReader struct {
	*csv.Reader
	HasHeader bool
}

type Option func(*csvReader)

type Record struct {
	Header []string
	Row    []string
}

type Writer interface {
	Begin() error
	End() error
	WriteRecord(*Record) error
}

func Convert(reader io.Reader, w Writer, opts ...Option) error {
	r := &csvReader{
		Reader: csv.NewReader(reader),
	}
	for _, opt := range opts {
		opt(r)
	}

	err := w.Begin()
	if err != nil {
		return fmt.Errorf("convert begin: %w", err)
	}

	var record Record
	for row := 0; ; row++ {
		record.Row, err = r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return fmt.Errorf("read row %d: %w", row, err)
		}

		if len(record.Header) == 0 {
			record.Header = make([]string, len(record.Row))
			if r.HasHeader {
				copy(record.Header, record.Row)
				continue
			} else {
				for i := range record.Row {
					record.Header[i] = strconv.Itoa(i)
				}
			}
		}

		err := w.WriteRecord(&record)
		if err != nil {
			return fmt.Errorf("write record (row %d): %w", row, err)
		}
	}

	err = w.End()
	if err != nil {
		return fmt.Errorf("convert end: %w", err)
	}

	return nil
}

func WithHeader() Option {
	return func(cr *csvReader) { cr.HasHeader = true }
}
