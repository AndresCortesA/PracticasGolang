package main

import (
	"context"
	"time"
)

func main() {
	/*
		un contexto puede derivar de un contexto padre, a la hora de la ejecucion del contexto,
		y cuando el contexto es cancelado o terminado, cualquier otro contexto que se derive de este
		tambien termina su ejecucion
	*/
	ctx := context.Background()
	/*
		contexto derivado del contexto padre `ctx`
		esta funcion nos devuelve dos resultados el primero es el contexto
		y el segundo es una funcion cancel con la cual trabajaremos para cancelar el contexto
	*/
	ctxCancel, cancel := context.WithCancel(ctx)
	//se toma como prioridad y se ejecuta
	cancel()

	//se ejecuta al final de todo
	// defer cancel()
	ctxTimeout, cancelTimeout := context.WithTimeout(ctxCancel, time.Second*2)
	//se ejecuta al final de todo
	defer cancelTimeout()

	select {
	case <-ctxTimeout.Done():
		println("terminado")
		// case <-ctxCancel.Done():
		// 	println("cancelado")
	}

}
