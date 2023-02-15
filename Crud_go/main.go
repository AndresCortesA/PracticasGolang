package main

import (
	"Crud_go/handlers"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	handlers.Index()
	handlers.Create()
	handlers.Insert()

	log.Println("Server running........................")
	http.ListenAndServe(":8080", nil)
}
