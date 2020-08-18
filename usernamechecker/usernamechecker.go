package main

import (
	"flag"
	"fmt"
	"gofullstack/validationkit"
)

const UsernameRegex string = `^@?(\w){1,15}$`

func main() {
	var usernameInput string
	flag.StringVar(&usernameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker")
	fmt.Println("Checking Syntax for Username, \""+usernameInput+"\", resulted in:", validationkit.CheckUsernameSyntax(usernameInput))
}
