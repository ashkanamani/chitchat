package main

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)
type Config struct {
	Address string `json:"Address"`
	ReadTimeout int `json:"ReadTimeout"`
	WriteTimeout int `json:"WriteTimeout"`
	Static string `json:"Static"`
}

const ConfigurationFile = "config.json"
var logger *log.Logger
func parseTemplateFiles(filenames ...string) (t *template.Template) {
	var files []string
	t = template.New("layout")
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}
	t = template.Must(t.ParseFiles(files...))
	return
}
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// for logging
// func info(args ...interface{}) {
// 	logger.SetPrefix("INFO ")
// 	logger.Println(args...)
// }

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// version
// func version() string {
// 	return "0.1"
// }
