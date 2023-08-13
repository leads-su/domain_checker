package checker

import (
	"crypto/tls"
	"fmt"
)

type CertChecker struct {
	Domain string
	Port   string
}

func CreateCertChecker(domain string, port string) *CertChecker {
	return &CertChecker{
		Domain: domain,
		Port:   port,
	}
}

func (cc *CertChecker) Do() error {
	_, err := tls.Dial(
		"tcp",
		fmt.Sprintf("%s:%s", cc.Domain, cc.Port),
		&tls.Config{
			InsecureSkipVerify: false,
		},
	)
	if err != nil {
		return fmt.Errorf("%s: %s", cc.Domain, err)
	}

	return nil
}
