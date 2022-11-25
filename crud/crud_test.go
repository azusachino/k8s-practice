package crud

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/azusachino/k8s/cfg"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCrud(t *testing.T) {
	crud()
}

func crud() {
	client := cfg.GetClient()

	namespace := "default"
	desired := corev1.ConfigMap{Data: map[string]string{"foo": "bar"}}
	desired.Namespace = namespace
	desired.GenerateName = "crud-practice-"
	labels := make(map[string]string)
	labels[corev1.LabelHostname] = "minikube"
	desired.Labels = labels

	// create the target configmap
	created, err := client.CoreV1().ConfigMaps(namespace).Create(
		context.Background(),
		&desired,
		metav1.CreateOptions{},
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Created ConfigMap %s/%s\n", namespace, created.GetName())

	if !reflect.DeepEqual(created.Data, desired.Data) {
		panic("Created ConfigMap has unexpected data")
	}

	time.Sleep(10 * time.Second)

	// Read the configmap
	read, err := client.CoreV1().ConfigMaps(namespace).Get(
		context.Background(),
		created.GetName(),
		metav1.GetOptions{},
	)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Read ConfigMap %s/%s\n", namespace, read.GetName())

	if !reflect.DeepEqual(read.Data, desired.Data) {
		panic("Read ConfigMap has unexpected data")
	}

	time.Sleep(10 * time.Second)

	// Update the configmap
	read.Data["foo"] = "updated"
	updated, err := client.CoreV1().ConfigMaps(namespace).Update(
		context.Background(),
		read,
		metav1.UpdateOptions{},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Updated ConfigMap %s/%s\n", namespace, updated.GetName())

	if !reflect.DeepEqual(updated.Data, read.Data) {
		panic("Updated ConfigMap has unexpected data")
	}

	time.Sleep(10 * time.Second)

	// Delete
	err = client.CoreV1().ConfigMaps(namespace).Delete(
		context.Background(),
		created.GetName(),
		metav1.DeleteOptions{},
	)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Deleted ConfigMap %s/%s\n", namespace, created.GetName())
}
