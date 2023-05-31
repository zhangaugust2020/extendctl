package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

// rootCmd 定义一下主体命令变量
var rootCmd = &cobra.Command{
	Use:   "extendctl",
	Args:  cobra.MinimumNArgs(1), //防止用户只使用kubectl命令导致卡死，这里设置必须要设置一个参数才可使用
	Short: "extendctl命令本体",
	Long:  `extendctl命令本体`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Kubeconfig 我们想要用kubectl 去调用K8S集群，需要携带集群的kubeconfig配置
var Kubeconfig, Namespace string

func init() {
	//利用主体变量去全局接收 携带的kubeconfig配置
	rootCmd.PersistentFlags().StringVar(&Kubeconfig, "kubeconfig", "", "$HOME/.kube/config")
}

// Initnode 接收上面获取参数的Kubeconfig变量，去判断是否有效
func Initnode() string {

	//下面的逻辑就是去查询用户有没有输入kubeconfig 的配置文件
	//如果没有输入，则默认去$HOME/.kube/config 去找认证文件，如果都没有则退出
	if Kubeconfig == "" {
		if home := os.Getenv("HOME"); home != "" {
			Kubeconfig = filepath.Join(home, ".kube", "config") //如果为空则填入我们设置好的路径
			Kubeconfig = filepath.ToSlash(Kubeconfig)           //将路径里面的斜杠转为正斜杠
			_, err := filepath.EvalSymlinks(Kubeconfig)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else {
				return Kubeconfig
			}
		} else {
			fmt.Println("没有找到kubeconfig文件、没有找到用户家目录，请确认操作系统是否为linux")
			os.Exit(1)
		}

	} else {
		//当判断用户输入文件路径后，查询该文件路径是否正确，如果不正确，则返回错误信息
		_, err := filepath.EvalSymlinks(Kubeconfig)
		if err != nil {
			fmt.Println("kubeconfig所指定的文件不存在，或格式有问题，程序退出")
			os.Exit(1)
		} else {
			return Kubeconfig
		}

	}
	return ""
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&Namespace, "namespace", "n", "default", "choose namespace")
	cobra.CheckErr(rootCmd.Execute()) //入口
}
