package checker

import "domain_checker/pkg/params"

type CertCheckerFactory struct {
	DefaultPort string
}

func CreateCertCheckerFactory() *CertCheckerFactory {
	return &CertCheckerFactory{
		DefaultPort: "443",
	}
}

func (factory *CertCheckerFactory) Make(domain string, params *params.LaunchParams) CheckerInterface {
	if factory.DefaultPort == "" {
		panic("default port for CertCheckerFactory not set")
	}
	return CreateCertChecker(domain, factory.DefaultPort)
}

//##########################################################################

func init() {
	GetCheckerRegistry().Add("cert", CreateCertCheckerFactory())
}
