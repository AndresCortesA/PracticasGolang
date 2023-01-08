package users

import (
	"fmt"
	"math/rand"
)

type User struct {
	Id   int
	Name string
}

// los structs en golang, tienen la capacidad de "heredar" de un struct anterior a esto se le llama Embeber un
// struct dentro de otro

type Employee struct {
	User
	Active bool
}

//INTERFACES------------------------------------------

type Cashier interface {
	CalcTotal(item ...float32) float32
	deactivate()
	reactivate()
}

type Admin interface {
	DesactivateEmployee(c Cashier)
}

//METODOS-------------------------------------

func NewEmployee(name string) *Employee {
	return &Employee{
		User: User{
			Id:   rand.Intn(1000),
			Name: name,
		},
		Active: true,
	}
}

func (e *Employee) CalcTotal(item ...float32) float32 {

	if !e.Active {
		fmt.Println("Cashier Deactivated")
		return 0
	}

	var sum float32

	for _, i := range item {
		sum += i
	}
	return sum * 1.15

}

func (e *Employee) deactivate() {
	e.Active = false
}

func (e *Employee) reactivate() {
	e.Active = true
}

func (e *Employee) DesactivateEmployee(c Cashier) {
	c.deactivate()
}
