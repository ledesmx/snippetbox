package main

import (
	"log"
	"net/http"
)

func main() {
	// Initialize a new servemux or router
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern
	// "/" is a catch-all regardless of their URL path
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	// start a new web server with ListenAndServe
	err := http.ListenAndServe(":4000", mux)
	// it returns an error that is always non-nil
	log.Fatal(err)
}
