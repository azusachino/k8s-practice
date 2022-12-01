package kubeconfig

import (
	"context"
	"fmt"
	"path"
	"testing"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func TestSimpleClient(t *testing.T) {

	// 模拟目标 Objects 的标签
	lbls := labels.Set{"node": "minikube", "app": "mysql"}

	selector := labels.NewSelector()

	// 选择器
	req, err := labels.NewRequirement("app", selection.Equals, []string{"mysql"})

	if err != nil {
		panic(err.Error())
	}
	selector = selector.Add(*req)

	if selector.Matches(lbls) {
		fmt.Printf("Selector %v matched label set %v\n", selector, lbls)
	} else {
		panic("Selector should have matched labels")
	}

	otherSelector, err := labels.Parse("app=mysql")
	if err != nil {
		panic(err.Error())
	}

	if otherSelector.Matches(lbls) {
		fmt.Printf("Selector %v matched label set %v\n", otherSelector, lbls)
	} else {
		panic("Selector should have matched labels")
	}

	// list pods
	home := homedir.HomeDir()
	config, _ := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	client, _ := kubernetes.NewForConfig(config)

	// access the api to list pods
	pods, _ := client.CoreV1().Pods("default").List(context.Background(), v1.ListOptions{})

	fmt.Printf("Pods count: %d\n", len(pods.Items))
}
