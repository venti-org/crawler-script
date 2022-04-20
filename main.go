package main

import (
	"fmt"
	"os"

	"github.com/venti-org/crawler-script/commands"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	verbose bool
)

func initConfig() {
	customFormatter := &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05.000",
		DisableTimestamp: false,
		DisableColors:    false,
	}

	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	if verbose {
		logrus.SetReportCaller(true)
	}
}

type Command = cobra.Command

func main() {
	cobra.OnInitialize(initConfig)

	rootCmd := &Command{
		Use:           "anti",
		Short:         "Anti is crawler framework",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "verbose log")

	rootCmd.AddCommand(
		commands.CrawlCmd,
		commands.VMCmd,
		commands.ExampleCmd,
	)

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
