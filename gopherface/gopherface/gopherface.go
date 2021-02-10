package main

import (
	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gofullstack/gopherface/handlers"
	"net/http"
	"os"
)

const WEBSERVERPORT = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/profile", handlers.MyProfileHandler).Methods("GET")
	r.HandleFunc("/profile/{username}", handlers.ProfileHandler).Methods("GET")

	http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))
	_ = http.ListenAndServe(WEBSERVERPORT, nil)
}
