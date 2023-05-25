package main

import (
	"html/template"
	"net/http"
	"github.com/ashkanamani/chitchat/data"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	_, err := session(writer, request)
	if err != nil {
		generateHTML(writer, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	files := []string{
		"template/layout.html",
		"template/navbar.html",
		"template/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))

	if threads, err := data.Threads(); err == nil {
		templates.Execute(writer, threads)

	}
}