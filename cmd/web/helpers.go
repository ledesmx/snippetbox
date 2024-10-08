package main

import (
	"log/slog"
	"net/http"
)

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
	)

	app.logger.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("uri", uri),
	)
	http.Error(
		w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError,
	)
}

func (app *Application) clientError(w http.ResponseWriter, code int) {
	http.Error(
		w,
		http.StatusText(code),
		code,
	)
}
