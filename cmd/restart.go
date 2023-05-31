package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"strconv"
	"time"
)

// podOwnDeployCmd represents the podOwnDeploy command
var restartCmd = &cobra.Command{
	Use:   "restart",
	Args:  cobra.ExactArgs(1),
	Short: "重启Deployment部署的Pod",
	Long:  `重启Deployment部署的Pod`,
	Run:   restart,
}

// restart 获取集群Deployment信息函数重启Pod
func restart(cmd *cobra.Command, args []string) {

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

	restartNum, _ := strconv.Atoi(deploy.ObjectMeta.Annotations["restarts"])
	restartNum = restartNum + 1
	deploy.ObjectMeta.Annotations["restarts"] = cast.ToString(restartNum)

	_, err = clientset.AppsV1().Deployments(Namespace).Update(context.TODO(), deploy, v1.UpdateOptions{})

	deploymentsClient := clientset.AppsV1().Deployments(Namespace)
	data := fmt.Sprintf(`{"spec": {"template": {"metadata": {"annotations": {"kubectl.kubernetes.io/restartedAt": "%s"}}}}}`, time.Now().Format("20060102150405"))
	_, err = deploymentsClient.Patch(context.TODO(), deploy.ObjectMeta.Name, types.StrategicMergePatchType, []byte(data), v1.PatchOptions{})

	fmt.Println(deploy.ObjectMeta.Name + "=======>" + cast.ToString(restartNum))

	if err != nil {
		log.Println("err ===> ", err)
	}

}

func init() {
	rootCmd.AddCommand(restartCmd)
}
