package c2j

import (
	"fmt"
)

type Format string

func (f *Format) String() string {
	return string(*f)
}

func (f *Format) Set(v string) error {
	switch v {
	case "json-array":
		*f = Format(v)
	case "jsonlines-array":
		*f = Format(v)
	case "jsonlines-object":
		*f = Format(v)
	case "json-object":
		*f = Format(v)
	default:
		return fmt.Errorf("unsupported format %s", v)
	}
	return nil
}

func (*Format) Type() string {
	return "format"
}
