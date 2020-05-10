package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()

	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch the URL, ", url, " and encountered this error: ", err)
	} else {
		fmt.Println("Contents of url, ", url, ", is:")
		contents, err := ioutil.ReadAll(response.Body)

		response.Body.Close()
		if err != nil {
			log.Fatal("Failed to read the response body and encountered this error: ", err)
		}
		fmt.Println(string(contents))
	}
}

func main() {
	var wg sync.WaitGroup
	var urlList = []string{
		"http://golang.org/",
		"http://www.google.com/",
		"http://www.youtube.com/",
	}
	for _, url := range urlList {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}
	wg.Wait()
}
