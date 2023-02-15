package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "index", nil)
	})
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "create", nil)
	})
	log.Println("Server running........................")
	http.ListenAndServe(":8080", nil)
}
