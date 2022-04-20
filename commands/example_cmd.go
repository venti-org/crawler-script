package commands

import (
	_ "embed"
	"fmt"
)

//go:embed examples/httpbin.js
var exampleScript string

var ExampleCmd = &Command{
	Use: "example",
	Run: func(cmd *Command, args []string) {
		fmt.Println(exampleScript)
	},
}
