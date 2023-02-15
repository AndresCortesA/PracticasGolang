package main

import (
	"CRUD_go/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()

	serverDoneChan := make(chan os.Signal, 1)
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	//Server
	srv := server.New(":8080")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()
	log.Println("SERVER RUN....................")

	<-serverDoneChan

	srv.Shutdown(ctx)
	log.Println("SERVER STOP....................")

}
