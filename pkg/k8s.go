package pkg

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8s struct {
	Client *kubernetes.Clientset
}

func NewK8s() (*K8s, error) {
	var config *rest.Config
	if _, err := strconv.ParseBool(os.Getenv("IN_CLUSTER_CONFIG")); err == nil {
		config, err = rest.InClusterConfig()

		if err != nil {
			return nil, err
		}
	} else {
		config, err = readKubeConfig()
		if err != nil {
			return nil, err
		}
	}
	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &K8s{
		Client: clientSet,
	}, nil
}

func readKubeConfig() (*rest.Config, error) {
	if home, err := os.UserHomeDir(); home != "" && err == nil {
		kubeConfig := flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
		if kubeConfig != nil {
			if err != nil {
				return nil, err
			}
		}

		config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
		if err != nil {
			return nil, err
		}
		return config, nil

	}

	return nil, fmt.Errorf("Can not read home dir")
}
