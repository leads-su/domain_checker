package main

import (
	"domain_checker/pkg/checker"
	"domain_checker/pkg/params"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// TODO:возможно перенести в params
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
func Thread(in chan string, lp *params.LaunchParams, checkers *[]checker.CheckerFactoryInterface) {
	for {
		domain, ok := <-in

		if !ok {
			break
		}

		// проход по списку фабрик череров и создание конкретного чекера для этого теста
		for _, c := range *checkers {
			err := c.Make(domain, lp).Do()

			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
}

func main() {
	checkerRegistry := checker.GetCheckerRegistry()
	lp := params.InitLaunchParams(checkerRegistry.GetAvailableNames())
	content, err := FileGetContents(lp.PathFile)

	if err != nil {
		panic(err)
	}

	in := make(chan string)

	// инициализация и заполнение списка фабрик чекеров
	var checkers []checker.CheckerFactoryInterface
	for _, checker := range lp.Checkers {
		checkers = append(checkers, checkerRegistry.Get(checker))
	}

	// создание горутин
	i := 0
	for i < lp.NumProc {
		i++
		go Thread(in, lp, &checkers)
	}

	// отправка списка доменов на проверку
	sites := strings.Split(content, "\n")
	for _, site := range sites {
		in <- site
	}
}
