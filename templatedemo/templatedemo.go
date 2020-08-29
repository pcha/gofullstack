package main

import (
	"html/template"
	"log"
	"net/http"
)

type Gopher struct {
	Name string
}

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	var gophername string
	gophername = r.URL.Query().Get("gophername")
	if gophername == "" {
		gophername = "Gopher"
	}
	gopher := Gopher{Name: gophername}
	renderTemplate(w, "./templates/greeting.html", gopher)
}

func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080", nil)
}
