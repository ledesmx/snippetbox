package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Define a new command-line flag
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{
		logger: slog.New(slog.NewTextHandler(os.Stdout, nil)),
	}

	app.logger.Info("starting server", slog.String("addr", *addr))

	mux := app.routes()

	// start a new web server with ListenAndServe
	err := http.ListenAndServe(*addr, mux)
	// it returns an error that is always non-nil
	app.logger.Error(err.Error())
	os.Exit(1)
}
