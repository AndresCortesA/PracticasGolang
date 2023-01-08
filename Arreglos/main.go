package main

import "fmt"

//SLICES
func main() {
	/* 	// var numbers[2] int //creamos un arreglo con valor cero
	   	numbers := []int{0, 2, 4, 5}
	   	//len(arreglo) este nos ayuda a determinar el tamaño del arreglo
	   	//range arreglo este nos dice el  rango en que se encuentra el arreglo
	   	//---------------------------------------Pruebitas---------------------
	   	// var isTrue bool // creamos una variable de tipo boleano para saber su valor cero

	   	// PRUEBA CON LEN
	   	for i := 0; i < len(numbers); i++ {
	   		//numbers = append(numbers, i) // esta funcion append es para poner un conjunto de datos dentro del arreglo+
	   		fmt.Print(numbers[i], " ")
	   	}
	   	fmt.Print("\n")
	   	//PRUEBA CON RANGE
	   	for i := range numbers {
	   		fmt.Print(numbers[i], " ")
	   	}
	   	fmt.Print("\n")

	   	for i, number := range numbers {
	   		fmt.Println("Index:", i, " value:", number)
	   	}
	   	//para omitir valores en GO se debe utilizar "_" guion bajo ejemplo
	   	for _, number := range numbers {
	   		fmt.Println("value:", number)
	   	}

	   	//ARREGLO SIN MODIFICAR
	   	fmt.Println(numbers)
	   	//	fmt.Println(isTrue)
	*/
	//condicional en arreglos
	fruits := []string{"Manzana", "Pera", "Banano", "Maracuya"}

	/* 	for _, fruit := range fruits {
		if fruit != "Banano" {
			continue // palabra reservada
		}
		fmt.Println("Fruit:", fruit)
	} */

	for _, fruit := range fruits {
		index := len(fruit) - 1
		letter := fruit[index:]
		//fmt.Println(letter)
		//[inicial : final] extraer los datos en el rango indicado, en este caso [6:] va a
		//extraer de la palabra Manzana la "a" que se encuentra en la posicion 6 ya que va hasta el final, y
		// es un arreglo de cadenas al
		// fin y al cabo

		if letter == "o" {
			continue
		}
		fmt.Println("Fruit:", fruit)

	}

}
