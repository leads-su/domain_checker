package checker

import "domain_checker/pkg/params"

type CheckerFactoryInterface interface {
	Make(domain string, params *params.LaunchParams) CheckerInterface
}
