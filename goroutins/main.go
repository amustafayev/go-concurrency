package main

import (
	"fmt"
	"sync"

	"github.com/goroutins/mutex"
)

func printSmt(s string, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println(s)
}

func main() {
	// goRoutins()
	// challange.Challange()

	// mutex.MutexExample()
	mutex.ComplexMutexExample()

}

func goRoutins() {

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"gamma",
		"delta",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epilon",
	}
	wg.Add(len(words))

	for i, el := range words {
		go printSmt(fmt.Sprintf("%d : %s", i, el), &wg)
	}

	wg.Add(2)
	go printSmt("Something print 1", &wg)

	// time.Sleep(1 * time.Second)

	printSmt("Something print 2", &wg)
}
