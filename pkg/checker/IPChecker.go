package checker

import (
	"fmt"
	"net"
)

type IPChecker struct {
	Domain string
	IPs    *PermittedIPs
}

func CreateIPChecker(domain string, ips *PermittedIPs) *IPChecker {
	return &IPChecker{
		Domain: domain,
		IPs:    ips,
	}
}

func (ipc *IPChecker) Do() error {
	ips, err := net.LookupIP(ipc.Domain)
	if err != nil {
		return fmt.Errorf("%s - could not get IPs: %v", ipc.Domain, err)
	}

	if len(ips) > 1 {
		return fmt.Errorf("%s - more then 1 ips", ipc.Domain)
	}

	ip := ips[0].String()
	if !ipc.IPs.Contains(ip) {
		return fmt.Errorf("%s - invalid ip", ipc.Domain)
	}

	return nil
}
