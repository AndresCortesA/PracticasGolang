package countries

import (
	"errors"
	"fmt"
)

var myCountry string

//se declaran en minusculas las clases private = solo se puede utilizar dentro del package
//se declaran en mayusculas las clases public
var collection []string

var errNotFound error = errors.New("COUNTRY NOT FOUND")

//la funcion Add sirve para añadir contenido a la lista collection
func Add(country string) {
	collection = append(collection, country)
}

//la funcion SetMyCountrie nos sirve para seleccionar nuestro pais favorito
func SetMyCountry(country string) error {
	if !isInCollection(country) {
		return errNotFound
	}

	myCountry = country
	return nil
}

//List muestra una Lista ordenada con mi pais favorito
func List() {
	for i, countries := range collection {
		myCountryMsg := ""
		if countries == myCountry {
			myCountryMsg = "[my favorite countrie]"
		}
		fmt.Printf("%d. %s %s\n", i+1, countries, myCountryMsg)
	}
}

//Esta funcion nos ayuda a saber si existe un error con la libreria error que posee GO
//esta  libreria la utilizamos en SetMyCountrie la cual devuelve un error o el pais en caso de encontrarlo o no
func isInCollection(country string) bool {
	for _, countries := range collection {
		if countries == country {
			return true
		}
	}

	return false
}
