package main

import (
	"fmt"
	"gofullstack/validationkit"
	"net/http"
)

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Gopher")
}

func checkUsernameSyntaxHandler(w http.ResponseWriter, r *http.Request) {
	u := r.URL.Query().Get("username")
	if u == "" {
		http.Error(w, "Username not provided", http.StatusInternalServerError)
		return
	}

	check := validationkit.CheckUsernameSyntax(u)
	fmt.Fprintf(w, "Syntax check result for %v is %v", u, check)
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.HandleFunc("/username-syntax-checker", checkUsernameSyntaxHandler)
	http.ListenAndServe(":8080", nil)
}
