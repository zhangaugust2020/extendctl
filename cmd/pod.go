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

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: "与get命令一起使用，获取k8s上的pod信息",
	Long:  `与get命令一起使用，获取k8s上的pod信息`,
	Run:   connectPod,
}

// connectPod 获取集群pod信息函数
func connectPod(cmd *cobra.Command, args []string) {

	//前面链接的语法是相同的，后面可以把这些放在一个函数中
	kubeconfig := Initnode()
	var Kconfig *string
	Kconfig = &kubeconfig

	config, err := clientcmd.BuildConfigFromFlags("", *Kconfig)
	errorsf(err)
	clientset, err := kubernetes.NewForConfig(config)
	errorsf(err)

	//通过clientset去获取default 下的pod信息
	pod, err := clientset.CoreV1().Pods(Namespace).List(context.TODO(), v1.ListOptions{})

	//抄作业
	pods := pod.Items

	//遍历所有的pod信息
	for _, ns := range pods {

		fmt.Println(
			ns.ObjectMeta.Name,                    //pod名称
			ns.ObjectMeta.Labels,                  //pod标签
			ns.ObjectMeta.Namespace,               //命名空间
			ns.Spec.Containers[0].Name,            //控制器名称
			ns.Spec.Containers[0].ImagePullPolicy, //镜像拉取策略
			ns.Spec.Containers[0].Image,           //镜像名称
		)

	}

	if err != nil {
		log.Println("err ===> ", err)
	}

}

func init() {
	getCmd.AddCommand(podCmd)
}
