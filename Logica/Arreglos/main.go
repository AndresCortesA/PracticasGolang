package main

import "fmt"

/*
crearemos una funcion que haga un recorrido de un arreglo en el cual saque dos numeros de la funcion
ejemplo:

un arreglo {5,4,3,2,1,0}
sus posiciones {0,1,2,3,4,5}
lo que espero que salga: {5,4} esta es la primera iteracion
						 {4,3}
						 {3,2}
						 {+,+}
*/

func main() {

	arreglo := []int{5, 4, 3, 2, 1, 0}
	fmt.Println("un arreglo: ", arreglo)
	sweep(arreglo)

}

func sweep(numeros []int) {
	var (
		contador      int = len(numeros)
		primerNumero  int = 0
		segundoNumero int = 1
	)

	//creamos un for que recorra el arreglo hasta que el segundo número llegue al final del mismo

	for segundoNumero < contador {
		var posicion1 int = numeros[primerNumero]
		var posicion2 int = numeros[segundoNumero]
		fmt.Println("iteracion pares: {", posicion1, posicion2, "}")

		primerNumero++
		segundoNumero++

	}

}
