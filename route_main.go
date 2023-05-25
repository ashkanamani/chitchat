package main

import (
	"html/template"
	"net/http"
	"github.com/ashkanamani/chitchat/data"
)


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