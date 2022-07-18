package misc

import "github.com/teamssix/cf/command"

func init() {
	Command.RootCmd.AddCommand(versionCmd)
	Command.RootCmd.AddCommand(upgradeCmd)
}
