package pkg

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"os"
	"sync"
	"time"

	v1 "k8s.io/api/core/v1"

	"github.com/ragoob/tls-notifer/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Tls struct {
	K8s       *K8s
	NameSpace string
}

func (t *Tls) GetSecertsList() (*v1.SecretList, error) {
	result, err := t.K8s.Client.CoreV1().Secrets(t.NameSpace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Error("Could not get Secerts in nameSpace [%v] [%v]", t.NameSpace, err.Error())
		return nil, err
	}

	return result, nil

}

func (t *Tls) Watch(wg *sync.WaitGroup, quit chan bool) {
	defer wg.Done()
	ticker := time.NewTicker(10 * time.Second)
	for {
		select {
		case <-ticker.C:
			result, err := t.GetSecertsList()
			if err != nil {
				log.Error(err)
			}

			for _, v := range result.Items {
				if v.Data["tls.crt"] != nil && v.Data["tls.key"] != nil {
					certificates := []types.Certificate{}
					output, err := ParsePEM(v.Data["tls.crt"])

					if err != nil {
						log.Error(err)
						continue
					}

					for _, c := range output {
						certificates = append(certificates, types.Certificate{
							NotAfter:  c.NotAfter,
							NotBefore: c.NotBefore,
							Domains:   c.DNSNames,
							Path:      fmt.Sprintf("Kube-%s/%s/%s", os.Getenv("CLUSTER_NAME"), t.NameSpace, v.Name),
							Issuer:    c.Issuer.CommonName,
							ExpireIn:  c.NotAfter.Sub(time.Now()).Hours() / 24,
						})
					}

					ex := Exporter{}
					ex.Export(certificates)

				}
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}
}
