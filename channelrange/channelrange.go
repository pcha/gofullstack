package main

import "fmt"

func main() {
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"

	close(messageQueue)

	for m := range messageQueue {
		fmt.Println(m)
	}
}
