package c2j_test

import (
	"io"
	"strings"
	"testing"

	"github.com/pastdev/csvjson/pkg/c2j"
	"github.com/stretchr/testify/require"
)

func TestConvert(t *testing.T) {
	tester := func(
		name string,
		csv string,
		opts []c2j.Option,
		writer func(io.Writer) c2j.Writer,
		expected string,
	) {
		t.Run(name, func(t *testing.T) {
			var w strings.Builder
			err := c2j.Convert(
				strings.NewReader(csv),
				writer(&w),
				opts...)
			require.NoError(t, err)
			require.Equal(t, expected, w.String())
		})
	}

	tester(
		"json object writer without header",
		`kermit,thefrog
miss,piggy`,
		nil,
		func(w io.Writer) c2j.Writer { return c2j.NewJSONObjectWriter(w) },
		`[{"0":"kermit","1":"thefrog"},{"0":"miss","1":"piggy"}]`)

	tester(
		"json object writer with header",
		`firstname,lastname
kermit,thefrog
miss,piggy`,
		[]c2j.Option{c2j.WithHeader()},
		func(w io.Writer) c2j.Writer { return c2j.NewJSONObjectWriter(w) },
		`[{"firstname":"kermit","lastname":"thefrog"},{"firstname":"miss","lastname":"piggy"}]`)

	tester(
		"jsonlines object writer without header",
		`kermit,thefrog
miss,piggy`,
		nil,
		func(w io.Writer) c2j.Writer { return c2j.NewJSONLinesObjectWriter(w) },
		`{"0":"kermit","1":"thefrog"}
{"0":"miss","1":"piggy"}
`)

	tester(
		"jsonlines object writer with header",
		`firstname,lastname
kermit,thefrog
miss,piggy`,
		[]c2j.Option{c2j.WithHeader()},
		func(w io.Writer) c2j.Writer { return c2j.NewJSONLinesObjectWriter(w) },
		`{"firstname":"kermit","lastname":"thefrog"}
{"firstname":"miss","lastname":"piggy"}
`)

	tester(
		"json array writer without header",
		`kermit,thefrog
miss,piggy`,
		nil,
		func(w io.Writer) c2j.Writer { return c2j.NewJSONArrayWriter(w) },
		`[["kermit","thefrog"],["miss","piggy"]]`)

	tester(
		"json array writer with header",
		`firstname,lastname
kermit,thefrog
miss,piggy`,
		[]c2j.Option{c2j.WithHeader()},
		func(w io.Writer) c2j.Writer { return c2j.NewJSONArrayWriter(w) },
		`[["kermit","thefrog"],["miss","piggy"]]`)

	tester(
		"jsonlines array writer without header",
		`kermit,thefrog
miss,piggy`,
		nil,
		func(w io.Writer) c2j.Writer { return c2j.NewJSONLinesArrayWriter(w) },
		`["kermit","thefrog"]
["miss","piggy"]
`)

	tester(
		"jsonlines array writer with header",
		`firstname,lastname
kermit,thefrog
miss,piggy`,
		[]c2j.Option{c2j.WithHeader()},
		func(w io.Writer) c2j.Writer { return c2j.NewJSONLinesArrayWriter(w) },
		`["kermit","thefrog"]
["miss","piggy"]
`)
}
