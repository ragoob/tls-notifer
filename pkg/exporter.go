package pkg

import (
	"fmt"
	"log"

	"github.com/ragoob/tls-notifer/types"
)

type Exporter struct {
}

func (e *Exporter) Export(certs []types.Certificate) {
	for _, v := range certs {
		log.Println(fmt.Sprintf("[%v] expire In [%v]", v.Path, v.ExpireIn))
	}
}
