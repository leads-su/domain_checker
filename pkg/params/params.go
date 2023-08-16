package params

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type LaunchParams struct {
	NumProc  int
	IPs      []string
	Checkers []string
	PathFile string
}

func InitLaunchParams(availableCheckers *[]string) *LaunchParams {
	// options

	numProc := 5
	flag.IntVar(&numProc, "n", 5, "count routines [1, 100], default 5")

	ips := ""
	flag.StringVar(&ips, "a", "", "list of IP for check A record in DNS")

	dryRun := false
	flag.BoolVar(&dryRun, "dry-run", false, "dry run :)")

	usageDesc := fmt.Sprintf(
		"Usage of %s:\n\n%s",
		os.Args[0],
		fmt.Sprintf("%s [OPTIONS] checker_list path_of_file", os.Args[0]),
	)
	requiredDesc := fmt.Sprintf(
		"%s\n%s",
		fmt.Sprintf("  checker_list - checker list, available: %s", strings.Join(*availableCheckers, ", ")),
		fmt.Sprintf("  path_of_file - path to file with domain divided a new lines"),
	)

	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"%s\n\n%s\n\n",
			usageDesc,
			requiredDesc,
		)
		flag.PrintDefaults()
	}
	flag.Parse()

	lp := &LaunchParams{}

	if len(flag.Args()) < 2 {
		lp.error("Not found required args: checker_list path_of_file")
	}

	checkers := flag.Args()[0]
	pathFile := flag.Args()[1]

	lp.SetNumProc(numProc)
	lp.SetIPs(ips)
	lp.SetCheckers(checkers, availableCheckers)
	lp.SetPathFile(pathFile)

	fmt.Printf("checkers: %s, threads: %d, file: %s\n", checkers, numProc, pathFile)

	// если сухой запуск то завершаем работу
	if dryRun {
		fmt.Println("This is dry run, everything is fine, you can try a real launch!")
		os.Exit(0)
	}

	return lp
}

func (lp *LaunchParams) SetNumProc(numProc int) *LaunchParams {
	if numProc < 1 || numProc > 100 {
		lp.error(fmt.Sprintf("unacceptable number of process [%d], should be [1, 100]", numProc))
	}

	lp.NumProc = numProc

	return lp
}

func (lp *LaunchParams) SetCheckers(checkers string, availableCheckers *[]string) *LaunchParams {
	a := strings.Split(checkers, ",")

	inArray := func(needle string, a *[]string) bool {
		if a == nil {
			return false
		}
		for _, key := range *a {
			if needle == key {
				return true
			}
		}
		return false
	}

	for _, checker := range a {
		if !inArray(checker, availableCheckers) {
			availableCheckers := strings.Join(*availableCheckers, ", ")
			lp.error(fmt.Sprintf("checker %s not found, available checkers: %s", checker, availableCheckers))
		}

		lp.Checkers = append(lp.Checkers, checker)
	}

	return lp
}

func (lp *LaunchParams) SetIPs(ips string) *LaunchParams {
	if len(ips) == 0 {
		return lp
	}

	a := strings.Split(ips, ",")
	pattern := `^\d+\.\d+\.\d+\.\d+$`

	for _, ip := range a {
		ip = strings.TrimSpace(ip)
		if len(ip) == 0 {
			continue
		}
		matched, _ := regexp.MatchString(`^\d+\.\d+\.\d+\.\d+$`, ip)

		if !matched {
			lp.error(fmt.Sprintf("one or more IP not matched pattern %s", pattern))
		} else {
			lp.IPs = append(lp.IPs, ip)
		}
	}

	return lp
}

func (lp *LaunchParams) SetPathFile(pathFile string) *LaunchParams {
	if _, err := os.Stat(pathFile); errors.Is(err, os.ErrNotExist) {
		lp.error(fmt.Sprintf("file %s not found", pathFile))
	}
	lp.PathFile = pathFile
	return lp
}

func (lp *LaunchParams) error(text string) {
	fmt.Println(text)
	flag.Usage()
	os.Exit(1)
}
