package crud

import (
	"context"
	"fmt"
	"testing"

	"github.com/azusachino/k8s/cfg"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestError(t *testing.T) {
	doError()
}

func doError() {
	client := cfg.GetClient()

	namespace := "default"

	// ERR_NOT_FOUND
	_, err := client.CoreV1().ConfigMaps(namespace).Get(
		context.Background(),
		"not_exist_config_map",
		metav1.GetOptions{},
	)

	if err == nil {
		panic("ERR_NOT_FOUND expected")
	}

	if !errors.IsNotFound(err) {
		panic(err.Error())
	}

	cm := corev1.ConfigMap{Data: map[string]string{"foo": "bar"}}
	cm.Namespace = namespace
	cm.GenerateName = "k8s-practice-prefix-"
	lbls := make(map[string]string)
	lbls[corev1.LabelHostname] = "minikube"
	cm.Labels = lbls

	cr, err := client.CoreV1().ConfigMaps(namespace).Create(
		context.Background(),
		&cm,
		metav1.CreateOptions{},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created ConfigMap %s/%s\n", namespace, cr.GetName())

	dup := corev1.ConfigMap{}
	dup.Namespace = namespace
	dup.Name = cr.Name
	_, err = client.CoreV1().ConfigMaps(namespace).Create(
		context.Background(),
		&dup,
		metav1.CreateOptions{},
	)
	if err == nil {
		panic("ERR_ALREADY_EXISTS expected")
	}
	if !errors.IsAlreadyExists(err) {
		panic(err.Error())
	}

	cr.Data["update"] = "true"
	ur, err := client.CoreV1().ConfigMaps(namespace).Update(
		context.Background(),
		cr,
		metav1.UpdateOptions{},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Updated ConfigMap %s/%s\n", namespace, ur.GetName())

	// ERR_CONFICT
	cr.Data["conflict"] = "true"
	_, err = client.CoreV1().ConfigMaps(namespace).Update(
		context.Background(),
		cr,
		metav1.UpdateOptions{},
	)
	if err == nil {
		panic("ERR_CONFLICT expected")
	}
	if !errors.IsConflict(err) {
		panic(err.Error())
	}

}
