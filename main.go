package main

import (
	"log"
	"net/http"
)

// Define a home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// Write a byte slice  as the response body
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	// Initialize a new servemux or router
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern
	// "/" is a catch-all regardless of their URL path
	mux.HandleFunc("/", home)

	log.Print("starting server on :4000")

	// start a new web server with ListenAndServe
	err := http.ListenAndServe(":4000", mux)
	// it returns an error that is always non-nil
	log.Fatal(err)
}
