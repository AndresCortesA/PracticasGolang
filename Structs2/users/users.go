package users

import "fmt"

type User struct {
	Id      int
	Name    string
	Age     int
	friends []User
}

//Esta funcion tiene un metodo recive e imprime el valor que se necita de la direccion en este caso el nombre
func (n User) SayHello() {
	fmt.Println("Hola me llamo: ", n.Name)
}

//Esta funcion para ser modificada necesita la direccion de usuario por lo tanto se usa un puntero *User
func (u *User) AddFriend(friend User) {
	u.friends = append(u.friends, friend)
}

//Funcion para poner en lista los amigos del usuario
func (u User) ListFriends() {
	for i, f := range u.friends {
		fmt.Printf("%d. %s\n", i+1, f.Name)
	}
}

//Esta funcion se encargara de limpiar un amigo del usuario que se asigne (recibe un Id de user)
func (u *User) RemoveFriend(Id int) {
	index := u.findFriend(Id)

	if index < 0 {
		return
	}

	u.friends = append(u.friends[:index], u.friends[index+1:]...)
}

//Esta funcion se encarga de buscar el Id de un User y devuelve un indice si se encuentra, en caso de que no retorna un
//indice negativo
func (u User) findFriend(Id int) int {
	for i, f := range u.friends {
		if f.Id == Id {
			return i
		}
	}
	return -1
}
