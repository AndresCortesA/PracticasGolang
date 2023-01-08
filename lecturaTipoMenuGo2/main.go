package main

import "fmt"

func main() {
	menu :=
		`
	Bienvenido
	1 - pizza
	2 - tacos
	3 - no te gusta ninguna
	4 - salir
	`

	fmt.Print(menu)

	var eleccion int // Declaramos esta variable tipo scanner, es obligatorio
	fmt.Scanln(&eleccion)

	switch eleccion {
	case 1:
		fmt.Println("Te gusta las pizzas")
	case 2:
		fmt.Println("Te gustan los tacos")
	case 3:
		fmt.Println("No prefieres ninguna")
	case 4:
		break

	}

}
