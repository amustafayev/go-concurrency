package sleepingbarber

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HairCutDuration time.Duration
	NumberOfBarbers int
	BarbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addBarber(barber string) {

	shop.NumberOfBarbers++

	go func() { // execute the job
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(shop.ClientsChan) == 0 { // All barbers listen same client chan
				isSleeping = true
				color.Yellow("barber %s takes a nap", barber)
			}

			client, shopDone := <-shop.ClientsChan

			if shopDone {
				if isSleeping {
					color.Yellow("clietn %s wakes up barber %s", client, barber)
					isSleeping = false
				}

				//cut hair
				shop.cutHair(barber, client)
			} else {
				//Client channel closed, means shop closed. Sent barber home
				shop.sendBarberHome(barber)
				return // end of the job for current/added barber
			}

		}

	}()

}

func (shop *BarberShop) cutHair(barber string, client string) {
	color.Cyan("Barber %s cuts client %s's hair!", barber, client)
	time.Sleep(shop.HairCutDuration)
	color.Cyan("Barber %s finished to cut client %s's hair!", barber, client)
}

func (shop *BarberShop) sendBarberHome(barber string) {
	color.Green("Barber %s done, so he is going home", barber)
	shop.BarbersDoneChan <- true
}

func (shop *BarberShop) ShopClose() {
	color.Green("Shop is closing for today!")

	close(shop.ClientsChan)
	shop.Open = false // No more client accept. Just finish what left

	//We should make sure we sent barbers home
	for a := 1; a <= shop.NumberOfBarbers; a++ {
		<-shop.BarbersDoneChan // blocks if there is no message size of NumberOfBarbers
	}
	close(shop.BarbersDoneChan)

	color.Green("---------------------------")
	color.Green("Barber Shop Done For Today!")
}

func (shop *BarberShop) addClient(client string) {
	// print out a message
	color.Green("*** %s arrives!", client)

	if shop.Open {
		select {
		case shop.ClientsChan <- client:
			color.Yellow("Client %s takes a seat!", client)
		default:
			color.Red("No places left. client %s left!", client)
		}
	} else {
		color.Red("Shop is already Closed. Return Next Day!")
	}

}
