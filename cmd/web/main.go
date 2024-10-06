package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func main() {
	// Define a new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &Application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	// Initialize a new servemux or router
	mux := http.NewServeMux()

	fileHandler := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileHandler))

	// register the home function as the handler for the "/" URL pattern
	// "/" is a catch-all regardless of their URL path
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	app.logger.Info("starting server", slog.String("addr", *addr))

	// start a new web server with ListenAndServe
	err := http.ListenAndServe(*addr, mux)
	// it returns an error that is always non-nil
	app.logger.Error(err.Error())
	os.Exit(1)
}
