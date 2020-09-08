package main

import (
	"github.com/gorilla/mux"
	"gofullstack/gopherface/handlers"
	"net/http"
)

const WEBSERVERPORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")

	http.Handle("/", r)
	_ = http.ListenAndServe(WEBSERVERPORT, nil)
}
