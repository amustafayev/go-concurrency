package channels

import (
	"fmt"
	"time"
)

func SelectExample() {

	channel1 := make(chan string)

	channel2 := make(chan string)

	go func(ch chan string) {
		for {
			time.Sleep(6 * time.Second)
			ch <- fmt.Sprintf("Message from server 1")
		}
	}(channel1)

	go func(ch chan string) {
		for {
			time.Sleep(3 * time.Second)
			ch <- fmt.Sprintf("Message from server 2")
		}
	}(channel2)

	for {
		select {
		case s1 := <-channel1:
			fmt.Printf("Message from one: %s\n", s1)

		case s1 := <-channel1:
			fmt.Printf("Message from two: %s\n", s1)

		case s2 := <-channel2:
			fmt.Printf("Message from tree: %s\n", s2)
		case s2 := <-channel2:
			fmt.Printf("Message from four: %s\n", s2)
		}
	}
}
