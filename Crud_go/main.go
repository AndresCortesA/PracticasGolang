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
	log.Println("Server running........................")
	http.ListenAndServe(":8080", nil)
}
