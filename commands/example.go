package commands

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute(args []string) {
	exampleCmd := newexampleCmd()
	exampleCmd.addCommands()
	exampleCmd.cmd.SetArgs(args)

	_, err := exampleCmd.cmd.ExecuteC()
	if err != nil {
		os.Exit(1)
	}
}

type exampleCmd struct {
	baseCmd
}

func newexampleCmd() *exampleCmd {
	cc := &exampleCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:   "example",
		Short: "example usage",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// show help message only
			cmd.Help()
		},
	}
	// cc.cmd.CompletionOptions = cobra.CompletionOptions{
	// 	HiddenDefaultCmd: true,
	// }
	cc.cmd.Version = getVersion()
	cc.cmd.SilenceUsage = true

	// cc.cmd.PersistentFlags().BoolP("debug", "", false, "enable debug output")

	return cc
}

func (cc *exampleCmd) getCommand() *cobra.Command {
	return cc.cmd
}

func (cc *exampleCmd) addCommands() {
	addCommands(
		cc.cmd,
		newVersionCmd(),
	)
}
