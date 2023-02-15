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
	handlers.Delete()
	handlers.Edit()
	handlers.Update()

	log.Println("Server running........................")
	http.ListenAndServe(":8080", nil)
}
