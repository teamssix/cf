package alibaba

import (
	"github.com/spf13/cobra"
	"github.com/teamssix/cf/command"
)

func init() {
	Command.RootCmd.AddCommand(alibabaCmd)
}

var alibabaCmd = &cobra.Command{
	Use:   "alibaba",
	Short: "执行与阿里云相关的操作 (Perform Alibaba Cloud related operations)",
	Long:  "执行与阿里云相关的操作 (Perform Alibaba Cloud related operations)",
}
