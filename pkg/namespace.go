package pkg

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NameSpace struct {
	K8s *K8s
}

func (n *NameSpace) GetNameSpacesList() (*v1.NamespaceList, error) {
	result, err := n.K8s.Client.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Printf("Could not get nameSpaces [%v]", err.Error())
		return nil, err
	}

	return result, nil
}
