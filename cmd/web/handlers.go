package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function
func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	}

	templates, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := templates.ExecuteTemplate(w, "base", nil); err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id_string := r.PathValue("id")
	id, err := strconv.Atoi(id_string)
	if err != nil || id <= 0 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display snippet for ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Display a form for creating a new snippet")
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Save a new snippet")
}
