package commands

import (
	"github.com/venti-org/crawler-core/engine"
	"github.com/venti-org/crawler-core/extensions"
	"github.com/venti-org/crawler-script/script"
)

var debugF bool
var extensionF bool

var CrawlCmd = &Command{
	Use:           "crawl [--stdin] [script_path]",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *Command, args []string) error {
		if body, err := GetStdinOrOneFile(cmd, args); err != nil {
			return err
		} else if body == nil {
			return nil
		} else {
			return runScriptSpider(string(body))
		}
	},
}

func init() {
	RegisterStdinOrOneFile(CrawlCmd)
	CrawlCmd.Flags().BoolVar(&debugF, "debug", false, "")
	CrawlCmd.Flags().BoolVar(&extensionF, "extension", false, "")
}

func runScriptSpider(source string) error {
	s, err := script.NewScriptSpiderWithScript(nil, source)
	if err != nil {
		return err
	}
	e, err := engine.NewEngineBuilder().WithSpider(s).AppendComponents(s).Build()
	if err != nil {
		return err
	}
	if extensionF {
		if se, err := script.NewScriptExtension(s.GetVM()); err != nil {
			return err
		} else {
			e.AddExtension(se)
		}
	}
	if debugF {
		e.AddExtension(extensions.NewLogExtension())
	}
	e.Run()
	return nil
}
