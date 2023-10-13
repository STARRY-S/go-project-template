package commands

import (
	"fmt"

	"github.com/STARRY-S/go-project-template/pkg/utils"
	"github.com/spf13/cobra"
)

type versionCmd struct {
	baseCmd
}

func newVersionCmd() *versionCmd {
	cc := &versionCmd{}

	cc.baseCmd.cmd = &cobra.Command{
		Use:     "version",
		Short:   "Show version",
		Example: "  example version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("  example version %s\n", getVersion())
		},
	}

	return cc
}

func (cc *versionCmd) getCommand() *cobra.Command {
	return cc.cmd
}

func getVersion() string {
	if utils.GitCommit != "" {
		return fmt.Sprintf("%s - %s", utils.Version, utils.GitCommit)
	}
	return utils.Version
}
