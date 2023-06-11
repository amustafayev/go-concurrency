package diningphilospohers

import (
	"fmt"
	"sync"
	"time"

	"github.com/fatih/color"
)

type Philospohers struct {
	name      string
	rightFork int
	leftFork  int
}

//philopher list
var philospohers = []Philospohers{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotel", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

//constants
var hunger = 3 // how many times does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

var finished sync.Mutex
var order []string

func DiningPhilospohers() {
	//print out welcome message
	color.Cyan("Dining Philosphers Problem")
	color.Cyan("--------------------------")
	color.Yellow("The table is empty.")

	// starting meal
	dine()

	//print out finished message
	color.Yellow("The table is empty.")
}

func dine() {

	// eatTime = 0 * time.Second
	// thinkTime = 0 * time.Second
	// sleepTime = 0 * time.Second

	wg := &sync.WaitGroup{}
	wg.Add(len(philospohers))

	seated := &sync.WaitGroup{}
	seated.Add(len(philospohers))

	// it is not actually holds fork number. It just holds
	// first, second ... mutexes in order to lock certain block
	var forks = make(map[int]*sync.Mutex)

	// forks is a map of all 5 forks
	for i := 0; i < len(philospohers); i++ {
		forks[i] = &sync.Mutex{}
	}

	//start the meal

	for i := 0; i < len(philospohers); i++ {
		// Fire of the go routins

		go diningProblem(philospohers[i], wg, forks, seated)

	}

	wg.Wait()

	fmt.Println(order)
}

func diningProblem(philospoher Philospohers, wg *sync.WaitGroup, forks map[int]*sync.Mutex, seated *sync.WaitGroup) {
	defer wg.Done()

	//Seat the philophoser at the table
	fmt.Printf("%s is seated at the table.\n", philospoher.name)
	seated.Done()
	// wait untill all for blocks reach here
	seated.Wait()

	for i := hunger; i > 0; i-- {
		//lock both forks

		//note:
		//mutex guards the critical section of code that manipulates the shared data.
		// multiple goroutines are trying to acquire the lock associated with the /same/ mutex variable

		//
		if philospoher.leftFork > philospoher.rightFork {
			forks[philospoher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philospoher.name)
			forks[philospoher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philospoher.name)
		} else {
			forks[philospoher.leftFork].Lock()
			fmt.Printf("\t%s takes the left fork\n", philospoher.name)
			forks[philospoher.rightFork].Lock()
			fmt.Printf("\t%s takes the right fork\n", philospoher.name)
		}

		fmt.Printf("\t%s has both forks and eating\n", philospoher.name)
		time.Sleep(eatTime)

		fmt.Printf("\t%s thingking\n", philospoher.name)
		time.Sleep(thinkTime)

		forks[philospoher.leftFork].Unlock()
		forks[philospoher.rightFork].Unlock()

		fmt.Printf("\t%s put downs the forks!\n", philospoher.name)

	}

	finished.Lock()
	order = append(order, philospoher.name)
	finished.Unlock()

	fmt.Printf("%s is satisfied!\n", philospoher.name)
	fmt.Printf("%s left the table!\n", philospoher.name)

}
