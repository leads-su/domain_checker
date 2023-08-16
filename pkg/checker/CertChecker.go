package checker

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	dialer := tls.Dialer{
		Config: &tls.Config{
			InsecureSkipVerify: false,
		},
	}
	conn, err := dialer.DialContext(
		ctx,
		"tcp",
		fmt.Sprintf("%s:%s", cc.Domain, cc.Port),
	)
	cancel()

	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err != nil {
		return fmt.Errorf("%s: %s", cc.Domain, err)
	}

	return nil
}
