package misc

import (
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/pkg/util"
	"github.com/teamssix/cf/pkg/util/cmdutil"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "更新 cf 到最新版本 (Update cf to the latest version)",
	Long:  "更新 cf 到最新版本 (Update cf to the latest version)",
	Run: func(cmd *cobra.Command, args []string) {
		cmdutil.Upgrade(util.GetCurrentVersion())
	},
}
