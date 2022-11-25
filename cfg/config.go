package cfg

import (
	"os"
	"path"

	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// GetConfig from local .kube
func GetConfig() *restclient.Config {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	config, err := clientcmd.BuildConfigFromFlags("", path.Join(home, ".kube/config"))
	if err != nil {
		panic(err.Error())
	}

	return config
}

// GetClient the default k8s client
func GetClient() *kubernetes.Clientset {
	client, err := kubernetes.NewForConfig(GetConfig())
	if err != nil {
		panic(err.Error())
	}
	return client
}
