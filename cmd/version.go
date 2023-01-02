package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)

// podOwnDeployCmd represents the podOwnDeploy command
var versionCmd = &cobra.Command{
	Use:   "version",
	Args:  cobra.ExactArgs(1),
	Short: "A brief description of your command",
	Long:  ``,
	Run:   version,
}

// 获取集群node信息函数
func version(cmd *cobra.Command, args []string) {

	//前面链接的语法是相同的，后面可以把这些放在一个函数中
	kubeconfig := Initnode()
	var Kconfig *string
	Kconfig = &kubeconfig

	config, err := clientcmd.BuildConfigFromFlags("", *Kconfig)
	errorsf(err)
	clientset, err := kubernetes.NewForConfig(config)
	errorsf(err)

	//通过clientset去获取default 下的deploy信息
	deploy, err := clientset.AppsV1().Deployments(Namespace).Get(context.TODO(), args[0], v1.GetOptions{})

	fmt.Println(deploy.ObjectMeta.Annotations["restarts"])

	if err != nil {
		log.Println("err ===> ", err)
	}

}

func init() {
	rootCmd.AddCommand(versionCmd)
}
