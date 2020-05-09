package main

import "fmt"

func main() {
	messageQueue := make(chan string, 3)
	messageQueue <- "one"
	messageQueue <- "two"
	messageQueue <- "three"

	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
	fmt.Println(<-messageQueue)
}
