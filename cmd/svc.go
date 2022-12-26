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

var svcCmd = &cobra.Command{
	Use:   "svc",
	Short: "",
	Long:  ``,
	Run:   connectSvc,
}

func init() {
	getCmd.AddCommand(svcCmd)
}

// 定义svc查询
func connectSvc(cmd *cobra.Command, args []string) {

	//前面链接的语法是相同的，后面可以把这些放在一个函数中
	kubeconfig := Initnode()
	var Kconfig *string
	Kconfig = &kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *Kconfig)
	errorsf(err)
	clientset, err := kubernetes.NewForConfig(config)
	errorsf(err)

	//通过clientset去获取default 下的pod信息
	svc, err := clientset.CoreV1().Services(Namespace).List(context.TODO(), v1.ListOptions{})

	//抄作业
	svcs := svc.Items

	//fmt.Println(svcs)

	//遍历所有的pod信息
	for _, ns := range svcs {

		fmt.Println(ns.ObjectMeta.Name, //svc名称
			ns.ObjectMeta.Labels, //svc标签
			ns.Spec.Type,         //svc类型
			ns.Spec.Ports[0].NodePort,
			ns.Spec.Ports[0].Port,
		)

	}

	if err != nil {
		log.Println("err ===> ", err)
	}

}
