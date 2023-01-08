package data

import (
	"fmt"
	"sync"
)

type Book struct {
	ID       int
	Title    string
	Finished bool
}

var books = []*Book{
	{1, "dune", false},
	{2, "El perfume", false},
	{3, "The world of ice and fire", false},
	{4, "Teoria de la noche", false},
	{5, "Blanca Olmedo", false},
	{6, "El Principito", false},
	{7, "100 años de soledad", false},
	{8, "El alqumista", false},
	{9, "El libro del cementerio", false},
	{10, "Maze runner", false},
}

func findBook(id int, m *sync.Mutex) (int, *Book) {
	index := -1
	var book *Book
	m.Lock()
	for i, b := range books {
		if b.ID == id {
			index = i
			book = b
		}

	}
	m.Unlock()
	return index, book
}

func FinishBook(id int, m *sync.Mutex) {
	i, book := findBook(id, m)
	if i < 0 {
		return
	}

	m.Lock() // protegemos este valor para que no sea modificado por otras gorutines
	book.Finished = true
	books[i] = book
	m.Unlock() //para liberarlo
	fmt.Printf("Finished Book: %s\n", book.Title)
}
