package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page!")
	fmt.Fprintln("Endpoint hit: home page")
}

func handleRequests() {
	http.HandleFunc
}
