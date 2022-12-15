package pkg

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func ParsePEM(data []byte) ([]*x509.Certificate, error) {
	output := []*x509.Certificate{}

	for {
		block, rest := pem.Decode(data)
		if block == nil {
			break
		}

		data = rest
		if block.Type != "CERTIFICATE" {
			continue
		}

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("tried to parse malformed x509 data, %s", err.Error())
		}

		output = append(output, cert)
	}

	return output, nil
}
