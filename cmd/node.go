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

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "",
	Long:  ``,
	Run:   connectnode, //当访问node资源时，直接调用connectnode函数即可
}

// 获取集群node信息函数
func connectnode(cmd *cobra.Command, args []string) {

	kubeconfig := Initnode()
	var Kconfig *string
	Kconfig = &kubeconfig

	//通过client-go包去链接K8集群
	//需要传入master地址或者kubeconfig的路径
	config, err := clientcmd.BuildConfigFromFlags("", *Kconfig)
	errorsf(err)

	//创建一个客户端链接
	clientset, err := kubernetes.NewForConfig(config) //链接k8的rest接口后，我们会得到一个内存地址的切片
	errorsf(err)                                      //切片的每个数据都是一个接口

	//通过clientset去获取node信息
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})

	//过滤Items下的数据,也就是3个节点的json信息
	nodess := nodes.Items

	fmt.Printf("\nThere are %d namespaces in cluster\n", len(nodess))

	for _, ns := range nodess { //将3个node节点的json切片遍历出来
		//ns中存放的就是单个node的详细信息
		fmt.Printf(
			"Name: %s,"+
				" Status: %s,"+
				" CreateTime: %s"+
				"\n",

			ns.ObjectMeta.Name, //将单个节点下的信息打印出来
			ns.Status.Addresses[0].Address,
			ns.CreationTimestamp)
	}
	if err != nil {
		log.Println("err ===> ", err)
	}

}

// 处理上面的一堆错误检查
func errorsf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func init() {
	getCmd.AddCommand(nodeCmd)
}
