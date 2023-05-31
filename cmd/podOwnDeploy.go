package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strings"
)

// podOwnDeployCmd represents the podOwnDeploy command
var podOwnDeployCmd = &cobra.Command{
	Use:   "podOwnDeploy",
	Args:  cobra.ExactArgs(1),
	Short: "查询启动的Pod属于那个Deployment",
	Long:  `查询启动的Pod属于那个Deployment`,
	Run:   findDeploy,
}

// findDeploy 获取集群Deployment信息函数
func findDeploy(cmd *cobra.Command, args []string) {

	//前面链接的语法是相同的，后面可以把这些放在一个函数中
	kubeconfig := Initnode()
	var Kconfig *string
	Kconfig = &kubeconfig

	config, err := clientcmd.BuildConfigFromFlags("", *Kconfig)
	errorsf(err)
	clientset, err := kubernetes.NewForConfig(config)
	errorsf(err)

	//通过clientset去获取default 下的deploy信息
	deploy, err := clientset.AppsV1().Deployments(Namespace).List(context.TODO(), v1.ListOptions{})

	arr := strings.Split(args[0], "-")

	//获取deployment列表
	deploys := deploy.Items

	for _, dep := range deploys {
		if dep.ObjectMeta.Name == arr[0] {
			fmt.Println(dep.ObjectMeta.Name)
		}
	}

	if err != nil {
		log.Println("err ===> ", err)
	}

}

func init() {
	rootCmd.AddCommand(podOwnDeployCmd)
}
