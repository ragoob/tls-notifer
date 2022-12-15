package pkg

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/ragoob/tls-notifer/types"
)

type Exporter struct {
}

func (e *Exporter) Export(certs []types.Certificate) {
	for _, v := range certs {
		if v.ExpireIn <= 0 {
			log.Warn(fmt.Sprintf("[%v] expired", v.Path))

		} else {
			log.Info(fmt.Sprintf("[%v] expire In [%v] days", v.Path, int(v.ExpireIn)))

		}
	}
}
