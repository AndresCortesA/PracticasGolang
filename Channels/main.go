package main

/*
supongamos que tenemos una gourutine en el channel A[] el cual quiere compartir su informacion
A[] INFO -> CHANNEL
ahora tenemos otra gourutine B[] la cual quiere conocer la informacion del channel
B[] CHANNEL -> INFO

asi se conparte la informacion de una gourutine a otra solo desde el channel, osea no se comparten la
misma informacion


*/
import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {
	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	//Cuantas gourutines vamos a esperar eso, se soluciona con wg.add()
	wg.Add(2)

	go generadorDeID(wg, IDsChan)
	go logIDS(wg, IDsChan)

	wg.Wait()
}

func generadorDeID(wg *sync.WaitGroup, idsChan chan string) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		//recibimos el valor del id hacia la variable idsChan con el indice de la iteracion y el id
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}
	close(idsChan)

	wg.Done()
}

func logIDS(wg *sync.WaitGroup, idsChan chan string) {
	//y aqui recibimos la variable idsChan hacia la variable id
	for id := range idsChan {
		fmt.Println(id)
	}

	wg.Done()
}
