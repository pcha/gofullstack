package main

import (
	"fmt"
	"gofullstack/socialmedia"
)

func main() {
	myPost := socialmedia.NewPost(
		"PabloC",
		socialmedia.MoodStateHappy,
		"Go is awesome!",
		"Check out the Go web site",
		"https://golang.org",
		"",
		"",
		[]string{"go", "golang", "programming language"},
	)
	fmt.Printf("myPost %+v\n", myPost)
}
