package greetingspackage

import "fmt"

func PrintGreetings() {
	fmt.Println("Public Greetings!!")
}

func printGreetingsUnexported() {
	fmt.Println("private greetings!!")
}
