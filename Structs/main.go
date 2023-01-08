package main

import "fmt"

type Criatura struct {
	Name string
}

func main() {
	// los structs funcionan como formularios en papel que podrian usar, por ejemplo para declarar una criatura
	//con su respectivo nombre

	c := Criatura{
		Name: "Sammy el tiburon",
	}

	fmt.Println(c.Name)
}
