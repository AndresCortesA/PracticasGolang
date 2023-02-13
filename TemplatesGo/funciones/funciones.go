package funciones

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
}

func Personas(nombrePersona string, edadPersona int) *Persona {
	return &Persona{
		Nombre: nombrePersona,
		Edad:   edadPersona,
	}
}

func CargarTemplate(nombreArchivo string, data interface{}) {
	t, err := template.New(nombreArchivo).ParseFiles("templates/" + nombreArchivo)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
