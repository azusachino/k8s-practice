package main

import (
	"fmt"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
)

func main() {
	// 模拟目标Objects的标签
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
}
