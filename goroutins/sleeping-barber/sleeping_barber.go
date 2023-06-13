// This is a simple demonstration of how to solve the Sleeping Barber dilemma, a classic computer science problem
// which illustrates the complexities that arise when there are multiple operating system processes. Here, we have
// a finite number of barbers, a finite number of seats in a waiting room, a fixed length of time the barbershop is
// open, and clients arriving at (roughly) regular intervals. When a barber has nothing to do, he or she checks the
// waiting room for new clients, and if one or more is there, a haircut takes place. Otherwise, the barber goes to
// sleep until a new client arrives. So the rules are as follows:
//
//		- if there are no customers, the barber falls asleep in the chair
//		- a customer must wake the barber if he is asleep
//		- if a customer arrives while the barber is working, the customer leaves if all chairs are occupied and
//		  sits in an empty chair if it's available
//		- when the barber finishes a haircut, he inspects the waiting room to see if there are any waiting customers
//		  and falls asleep if there are none
// 		- shop can stop accepting new clients at closing time, but the barbers cannot leave until the waiting room is
//	      empty
//		- after the shop is closed and there are no clients left in the waiting area, the barber
//		  goes home
//
// The Sleeping Barber was originally proposed in 1965 by computer science pioneer Edsger Dijkstra.
//
// The point of this problem, and its solution, was to make it clear that in a lot of cases, the use of
// semaphores (mutexes) is not needed.

package sleepingbarber

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var NumberOfSeats = 10
var hairCutDuration = 1000 * time.Millisecond
var shopOpenInterval = 10

var arrivalRate = 100
var timeOpen = 10 * time.Second

func SleepingBarberProblem() {

	rand.Seed(time.Now().UnixNano())

	// Create barber shop
	barbersDone := make(chan bool)
	clientChannel := make(chan string, NumberOfSeats)

	shop := BarberShop{
		HairCutDuration: hairCutDuration,
		ShopCapacity:    NumberOfSeats,
		NumberOfBarbers: 0,
		BarbersDoneChan: barbersDone,
		ClientsChan:     clientChannel,
		Open:            true,
	}

	//add barber

	shop.addBarber("Federik")
	shop.addBarber("Jonny")
	shop.addBarber("Max")
	shop.addBarber("Dennis")
	shop.addBarber("Dazy")
	shop.addBarber("Alex")

	// add barber shop goroutins
	shopClosing := make(chan bool)
	closed := make(chan bool)

	//Closing function
	go func() {
		<-time.After(timeOpen)
		shopClosing <- true
		shop.ShopClose()
		closed <- true
	}()

	//add client
	go func() {

		i := 1

		for {
			randomMillseconds := rand.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				color.Red("Shop is closed! Client #%d leave", i)
				return
			case <-time.After(time.Millisecond * time.Duration(randomMillseconds)):
				shop.addClient(fmt.Sprintf("Client #%d", i))
				i++
			}
		}
	}()

	//blcok main goroutin untill barber closed

	<-closed

	// // time.Sleep(4 * time.Second)
	// // close(clientChannel)
	// time.Sleep(10 * time.Second)

}
