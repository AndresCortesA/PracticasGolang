package main

import (
	"TemplatesGo/funciones"
	_ "html/template"
)

func main() {

	andres := funciones.Personas("Andres", 20, []string{"<class>", "Programar", "Llorar por no encontrar los bugs"})
	funciones.CargarTemplate("index.html", andres)

}
