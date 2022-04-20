package commands

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

type Command = cobra.Command

var (
	fromStdin bool
)

func RegisterStdinOrOneFile(cmd *Command) {
	cmd.Flags().BoolVar(&fromStdin, "stdin", false, "")
}

func GetStdinOrOneFile(cmd *Command, args []string) ([]byte, error) {
	var reader io.Reader
	if fromStdin {
		reader = os.Stdin
	} else if len(args) == 1 {
		if f, err := os.Open(args[0]); err != nil {
			return nil, err
		} else {
			defer f.Close()
			reader = f
		}
	} else {
		return nil, cmd.Help()
	}
	return ioutil.ReadAll(reader)
}
