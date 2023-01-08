package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//fmt.Println("Hola mundo")
	menu :=
		`
	Bienvenido
	[ 1 ] Pizza
	[ 2 ] Tacos
	¿Qué prefieres?
		` // acento grave alt + 96 diferente de comillas simples

	fmt.Print(menu)

	reader := bufio.NewReader(os.Stdin)

	entrada, _ := reader.ReadString('\n')          // Leer hasta el separador de salto de línea
	eleccion := strings.TrimRight(entrada, "\r\n") //remover el salto de línea de la entrada del

	switch eleccion {
	case "1":
		fmt.Println("Prefires pizza")
	case "2":
		fmt.Println("Prefieres tacos")
	default:
		fmt.Println("No prefieres ninguno de ellos")

	}

}
