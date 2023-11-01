package root

import (
	"github.com/pastdev/csvjson/cmd/csvjson/c2j"
	"github.com/pastdev/csvjson/cmd/csvjson/version"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := cobra.Command{
		Use:   "csvjson",
		Short: `A tool for converting between CSV and JSON.`,
	}

	cmd.AddCommand(c2j.New())
	cmd.AddCommand(version.New())

	return &cmd
}
