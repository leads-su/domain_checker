package main

import (
	"domain_checker/pkg/checker"
	"domain_checker/pkg/params"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

func FileGetContents(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(b), nil
}

// функция логики потока
func Thread(wg *sync.WaitGroup, in chan string, lp *params.LaunchParams, checkers *[]checker.CheckerFactoryInterface) {
	defer wg.Done()
	for domain := range in {
		// проход по списку фабрик чекеров и создание конкретного чекера для этого теста
		for _, c := range *checkers {
			err := c.Make(domain, lp).Do()

			if err != nil {
				log.Printf("%v\n", err)
			}
		}
	}
}

func main() {
	log.SetFlags(0)
	checkerRegistry := checker.GetCheckerRegistry()
	lp := params.InitLaunchParams(checkerRegistry.GetAvailableNames())
	content, err := FileGetContents(lp.PathFile)

	if err != nil {
		panic(err)
	}

	in := make(chan string, 10)

	// инициализация и заполнение списка фабрик чекеров
	var checkers []checker.CheckerFactoryInterface
	for _, checker := range lp.Checkers {
		checkers = append(checkers, checkerRegistry.Get(checker))
	}

	var wg sync.WaitGroup
	wg.Add(lp.NumProc)

	// создание горутин
	i := 0
	for i < lp.NumProc {
		i++
		go Thread(&wg, in, lp, &checkers)
	}

	// отправка списка доменов на проверку
	sites := strings.Split(content, "\n")
	for _, site := range sites {
		in <- site
	}

	close(in)
	wg.Wait()
}
