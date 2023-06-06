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
