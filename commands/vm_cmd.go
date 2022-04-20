package commands

import (
	"github.com/venti-org/crawler-script/script"
)

var VMCmd = &Command{
	Use:           "vm [--stdin] [script_path]",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *Command, args []string) error {
		if body, err := GetStdinOrOneFile(cmd, args); err != nil {
			return err
		} else if body == nil {
			return nil
		} else {
			return runScript(string(body))
		}
	},
}

func init() {
	RegisterStdinOrOneFile(VMCmd)
}

func runScript(body string) error {
	vm := script.NewVM()
	_, err := vm.Run(body)
	return err
}
