package types

import (
	"time"
)

type Certificate struct {
	Domains   []string  `json:"domains"`
	NotAfter  time.Time `json:"notAfter"`
	NotBefore time.Time `json:"notBefore"`
	Path      string    `json:"path"`
	Issuer    string    `json:"issuer"`
	ExpireIn  string    `json:"expireIn"`
}
