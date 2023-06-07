package mutex

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMsg(s string, mx *sync.Mutex) {
	defer wg.Done()
	mx.Lock()
	msg = s
	mx.Unlock()
}

func printMessage() {
	fmt.Println(msg)
}

func MutexExample() {

	msg = "Hello initial"
	var mx sync.Mutex

	wg.Add(2)

	go updateMsg("Hello 1 ", &mx)
	go updateMsg("Hello 2 ", &mx)
	wg.Wait()
	printMessage()

	wg.Wait()
}

type Income struct {
	Source string
	Amount int
}

func ComplexMutexExample() {
	var bankBalance int

	fmt.Printf("Initial bank balance: %d.00", bankBalance)
	fmt.Println()

	var bank sync.Mutex

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investment", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			for week := 1; week <= 52; week++ {
				bank.Lock()         // do not let other threads acces
				temp := bankBalance // read OPERATION
				temp += income.Amount
				bankBalance = temp // write OPERATION
				bank.Unlock()
				fmt.Printf("On week %d, you earnd %d.00 from %s\n", week, income.Amount, income.Source)
			}

		}(i, income)

	}

	wg.Wait()
	fmt.Printf("Final bank balance: %d.00", bankBalance)

}
