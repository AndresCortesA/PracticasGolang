package main

import (
	"Goruitines/data"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Gorutines = parecido a un hilo pero su espacio en memoria dinámico manejados por GOLANG
//(eficaz para el uso de memoria) no son procesos en paralelo,
//son procesos concurrentes(se ejecutan una después de la otra) en cortos periodos de tiempo
//concurrencia = ejecutar procesos en cortos periodos a la vez

//Threads = hilos con espacio en memoria fijo manejados por el SO

//WaitGroups

//pruebas en comandos como go run --race main.go (esto nos da una data race de advertencia, que es si una gorutine
//quiere modificar el valor de una gorutine esta tratando de acceder), para arreglar esto vamos a usar Mutex en data.go

func main() {
	start := time.Now() // tiempo de ejecución
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		//	go showGorutines(i, wg)

		go readBook(i, wg, m)
	}
	//time.Sleep(time.Minute) //WaitGroups nos ayuda a monitoriar las rutines
	wg.Wait()                                    // WaitGroups = espera a que las gorutines terminen de ejecutarce
	duration := time.Since(start).Milliseconds() //finalización del tiempo de ejecucion

	fmt.Printf("duracion de las gorutines %dms\n", duration)
}

/* func showGorutines(id int, wg *sync.WaitGroup) {

	delay := rand.Intn(500)

	fmt.Printf("Gouruitine # %d with %dms\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done() // se encarga de terminar la ejecucion cuando todas las gorutines
} */

func readBook(id int, wg *sync.WaitGroup, m *sync.Mutex) {

	data.FinishBook(id, m)

	delay := rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(delay))
	wg.Done()
}
