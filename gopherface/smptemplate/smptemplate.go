package main

import (
	"fmt"
	"gofullstack/gopherface/socialmedia"
	"gofullstack/gopherface/templatesUtil"
	"net/http"
)

func displaySocialMediaPostHandler(w http.ResponseWriter, r *http.Request) {
	myPost := socialmedia.NewPost("pablocha", socialmedia.Moods["thrilled"], "Go is awesome!",
		"Check out the Go website", "https://golang.org", "/images/gogopher.png",
		"", []string{"go", "golang", "programming language"})
	fmt.Printf("myPost: %+v\n", myPost)
	templatesUtil.RenderTemplate(w, "./templates/socialmediapost.html", myPost)
}

func main() {
	http.HandleFunc("/display-social-media-post", displaySocialMediaPostHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":8080", nil)
}
