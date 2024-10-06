package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Define a new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Initialize a new servemux or router
	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileHandler))

	// register the home function as the handler for the "/" URL pattern
	// "/" is a catch-all regardless of their URL path
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("starting server", slog.String("addr", *addr))

	// start a new web server with ListenAndServe
	err := http.ListenAndServe(*addr, mux)
	// it returns an error that is always non-nil
	logger.Error(err.Error())
	os.Exit(1)
}
