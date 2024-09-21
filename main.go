package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function
func home(w http.ResponseWriter, r *http.Request) {
	// Write a byte slice  as the response body
	w.Write([]byte("Display the home page"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprintf("Display snippet for ID %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Save a new snippet"))
}

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
