package channels

import (
	"fmt"
	"strings"
)

func ChannelTest() {

	ping := make(chan string)
	pong := make(chan string)

	for {
		fmt.Print("Enter input-> ")

		var userInput string

		go func(ping <-chan string, pong chan<- string) {
			fmt.Print("\nwaiting")
			input := <-ping // function completely wait untill somthing sent to channel
			fmt.Println("Got input")
			pong <- fmt.Sprintf("%s!!!", strings.ToUpper(input))

		}(ping, pong)

		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			break
		}

		ping <- userInput
		fmt.Println(<-pong)
	}

	close(ping)
	close(pong)
}
