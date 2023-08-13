package checker

import "domain_checker/pkg/params"

type IPCheckerFactory struct {
	IPs *PermittedIPs
}

func CreateIPCheckerFactory() *IPCheckerFactory {
	return &IPCheckerFactory{
		IPs: CreatePermittedIps(),
	}
}

func (ipc *IPCheckerFactory) Make(domain string, params *params.LaunchParams) CheckerInterface {
	for _, ip := range params.IPs {
		ipc.IPs.Add(ip)
	}

	if ipc.IPs.Count() == 0 {
		panic("not set permitted ips for IPCheckerFactory")
	}
	return CreateIPChecker(domain, ipc.IPs)
}

//##########################################################################

func init() {
	GetCheckerRegistry().Add("ip", CreateIPCheckerFactory())
}
