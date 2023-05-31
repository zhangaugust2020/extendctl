package cmd

import (
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MinimumNArgs(1), //添加
	Short: "获取K8s上的数据信息，搭配其他命令一起使用",
	Long:  `获取K8s上的数据信息，搭配其他命令一起使用`,
	Run: func(cmd *cobra.Command, args []string) { //添加

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
