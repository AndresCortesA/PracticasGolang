package funciones

import (
	_ "html/template"
	"os"
	"text/template"
)

type Persona struct {
	Nombre  string
	Edad    int
	Hobbies []string
}

var incremento = template.FuncMap{
	"incremento": func(num int) int {
		return num + 1
	},
}

func Personas(nombrePersona string, edadPersona int, pasaTiempos []string) *Persona {
	return &Persona{
		Nombre:  nombrePersona,
		Edad:    edadPersona,
		Hobbies: pasaTiempos,
	}
}

func CargarTemplate(nombreArchivo string, data interface{}) {
	t, err := template.New(nombreArchivo).Funcs(incremento).ParseFiles("templates/" + nombreArchivo)
	if err != nil {
		panic(err)
	}

	err = t.Execute(os.Stdout, data)

	if err != nil {
		panic(err)
	}
}
