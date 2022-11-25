package kubeconfig

import (
	"fmt"
	"os"
	"path"
	"testing"

	"k8s.io/client-go/discovery"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
)

func TestYaml(t *testing.T) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}

	// 可被重复调用的 配置Getter
	kubeConfigGetter := func() (*api.Config, error) {
		kubeconfigYaml, err := os.ReadFile(path.Join(home, ".kube/config"))
		if err != nil {
			return nil, err
		}
		return clientcmd.Load([]byte(kubeconfigYaml))
	}

	config, err := clientcmd.BuildConfigFromKubeconfigGetter("", kubeConfigGetter)

	if err != nil {
		panic(err.Error())
	}

	client, err := discovery.NewDiscoveryClientForConfig(config)

	if err != nil {
		panic(err.Error())
	}

	ver, err := client.ServerVersion()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ver.String())
}
