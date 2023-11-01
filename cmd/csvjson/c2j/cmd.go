package c2j

import (
	"fmt"
	"os"

	"github.com/pastdev/csvjson/pkg/c2j"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	format := Format("jsonlines-object")
	var hasHeader bool

	cmd := cobra.Command{
		Use:   "c2j",
		Short: `Convert CSV to JSON`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var writer c2j.Writer

			switch format {
			case "json-array":
				writer = c2j.NewJSONArrayWriter(os.Stdout)
			case "jsonlines-array":
				writer = c2j.NewJSONLinesArrayWriter(os.Stdout)
			case "jsonlines-object":
				writer = c2j.NewJSONLinesObjectWriter(os.Stdout)
			case "json-object":
				writer = c2j.NewJSONObjectWriter(os.Stdout)
			}

			opts := []c2j.Option{}
			if hasHeader {
				opts = append(opts, c2j.WithHeader())
			}

			err := c2j.Convert(os.Stdin, writer, opts...)
			if err != nil {
				return fmt.Errorf("convert: %w", err)
			}

			return nil
		},
	}

	cmd.Flags().Var(&format, "format", "Format of JSON")
	cmd.Flags().BoolVar(
		&hasHeader,
		"has-header",
		false,
		"first line of CSV is a header")

	return &cmd
}
