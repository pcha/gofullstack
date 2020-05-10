package main

import (
	"fmt"
	"sync"
)

var greetings string
var howdyDone chan bool
var mutex = &sync.Mutex{}

func howdyGretings() {
	mutex.Lock()
	greetings = "Howdy Gopher!"
	mutex.Unlock()
	howdyDone <- true
}

func main() {
	howdyDone = make(chan bool, 1)
	go howdyGretings()
	mutex.Lock()
	greetings = "Hello Gopher!"
	fmt.Println(greetings)
	mutex.Unlock()
	<-howdyDone
}
