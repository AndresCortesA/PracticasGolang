package main

import (
	"Structs2/users"
	"fmt"
)

func main() {
	//Esta funcion va a utilizar la estructura(formulario) de tipo usuario que tendra; ID,name,edad
	martha := users.User{
		Id:   2430,
		Name: "Martha",
		Age:  28,
	}

	pedro := users.User{Id: 2145, Name: "Pedro", Age: 21}
	jhon := users.User{Id: 2657, Name: "Jhon", Age: 25}
	jane := users.User{Id: 2954, Name: "Jane", Age: 23}

	//existen formas para imprimir los valores de un usuario por ejemplo "%+v" esto nos devuelve el valor al que
	//pertenece cada valor que le asignemos al struct del usuario
	//fmt.Printf("%+v", martha)

	martha.SayHello()
	martha.AddFriend(pedro)
	martha.AddFriend(jhon)
	martha.AddFriend(jane)
	martha.ListFriends()
	martha.RemoveFriend(pedro.Id)
	fmt.Println("=====================")
	martha.ListFriends()

	//fmt.Printf("ID de usuario: %d\nNombre de usuario: %s\nEdad del usuario: %d\n", martha.ID, martha.Name, martha.Age)
}
