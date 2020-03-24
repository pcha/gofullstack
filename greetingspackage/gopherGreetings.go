package greetingspackage

import "fmt"

var MagicNumber int

func GopherGreetings() {
	fmt.Println("Hello gophers!")
	printGreetingsUnexported()
}

func init() {
	MagicNumber = 108
}
