package main

import (
	"sync"

	"github.com/ragoob/tls-notifer/pkg"
)

var wg sync.WaitGroup

func main() {

	k8s, err := pkg.NewK8s()
	if err != nil {
		panic(err)
	}

	ns := &pkg.NameSpace{
		K8s: k8s,
	}

	list, err := ns.GetNameSpacesList()
	if err != nil {
		panic(err)
	}
	quit := make(chan bool)
	wg.Add(len(list.Items))
	for _, v := range list.Items {
		tls := pkg.Tls{
			K8s:       k8s,
			NameSpace: v.Name,
		}

		go tls.Watch(&wg, quit)
	}

	wg.Wait()
}
