package main

import (
	"PatronDisenioSemaforo/designSemaforo"
	"time"
)

func main() {
	entradas, fuera := 1, 10*time.Second

	semaforo := designSemaforo.Nuevo(entradas, fuera)

	if err := semaforo.Adquirir(); err != nil {
		panic(err.Error())
	}

	if err := semaforo.Relacionar(); err != nil {
		panic(err.Error())
	}

}
