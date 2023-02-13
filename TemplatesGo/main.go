package main

import "TemplatesGo/funciones"

func main() {

	andres := funciones.Personas("Andres", 25)
	funciones.CargarTemplate("saludo.txt", andres)

}
