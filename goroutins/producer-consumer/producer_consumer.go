package producerconsumer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var total, pizzasMade, pizzasFailed int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	println(ch)
	p.quit <- ch
	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1

		fmt.Printf("\n\nReceived order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""

		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making Pizza #%d. It will take %d seconds.....\n", pizzaNumber, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("\n*** We run out of ingredients for pizza #%d!\n", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("\n*** The cook quit while making pizza #%d!\n", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("\nPizza ready order number #%d!\n", pizzaNumber)
		}

		p := PizzaOrder{
			message:     msg,
			success:     success,
			pizzaNumber: pizzaNumber,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzeria(pizzaMaker *Producer) {

	var i = 0

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			//we made pizza and sent it to producer channel
			case pizzaMaker.data <- *currentPizza:
			case quitChan := <-pizzaMaker.quit: // Only works when somthing sent to quit channel
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}

}

func ProducerConsumerEx() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	//print out message
	color.Cyan("The Pizzaria is open for business!")
	color.Cyan("----------------------------------")

	//create a producer
	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	//run producer in the background
	go pizzeria(pizzaJob)

	//create and run consumer
	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {

			if i.success {
				color.Green(i.message)
				color.Green("[CONSUMER]Order number #%d is out for delivery\n", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("[CONSUMER]The customer is really mad!\n")
			}
		} else {
			color.Cyan("[CONSUMER]----------------------")
			color.Cyan("[CONSUMER]Done making pizzas...")
			err := pizzaJob.Close()
			if err != nil {
				color.Red("[CONSUMER]" + "****Error closing channel")
			}
		}

	}

	// print the ending message

	color.Yellow("\n\n----------------")
	color.Yellow("Done For Today")

	switch {
	case pizzasFailed >= 9:
		color.Red("Today is terrible day")
	case pizzasFailed >= 6:
		color.Red("Today is a bad day")
	case pizzasFailed >= 4:
		color.Yellow("It was okay...")
	case pizzasFailed >= 2:
		color.Green("It was a pretty good day!")
	default:
		color.Green("It was a great day!")
	}
}

/**
Quit signal sent just by closing channel
*/
func CloseSimulation() {

	quitChan := make(chan chan error)

	go func(ch chan chan error) {
		select {
		case _ = <-ch:
			fmt.Println("Channel closed triggered")
		}

	}(quitChan)

	time.Sleep(5 * time.Second)

	close(quitChan)
	time.Sleep(5 * time.Second)

}
