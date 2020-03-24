package main

import (
	"fmt"
	"gofullstack/greetingspackage"
)

func main() {
	greetingspackage.PrintGreetings()
	greetingspackage.GopherGreetings()

	fmt.Println("The value of the Magic Number is:", greetingspackage.MagicNumber)
}
