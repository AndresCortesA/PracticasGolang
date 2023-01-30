package main

/*
supongamos que tenemos una gorutine en el channel A[] el cual quiere compartir su informacion
A[] INFO -> CHANNEL
ahora tenemos otra gorutine B[] la cual quiere conocer la informacion del channel
B[] CHANNEL -> INFO

asi se conparte la informacion de una gorutine a otra solo desde el channel, osea no se comparten la
misma informacion

tipos de channels:
bidireccionales - envia y recibe datos
receive-only - solo recibe
send-only solo envia

sin embargo podemos indicarle a cada channel que tendra un comportamiento en especifico

*/
import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

func main() {
	wg := &sync.WaitGroup{}
	IDsChan := make(chan string)
	fakeIdsChan := make(chan string)
	closedChans := make(chan int)
	//Cuantas gourutines vamos a esperar eso, se soluciona con wg.add()
	wg.Add(3)

	go generadorDeID(wg, IDsChan, closedChans)
	go generadorFakeID(wg, fakeIdsChan, closedChans)

	go logIDS(wg, IDsChan, fakeIdsChan, closedChans)

	wg.Wait()
}

func generadorFakeID(wg *sync.WaitGroup, fakeIdsChan chan<- string, closedChans chan<- int) {
	for i := 0; i < 50; i++ {
		id := uuid.New()
		fakeIdsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}
	close(fakeIdsChan)
	closedChans <- 1
	wg.Done()
}

//solo recibe <-
func generadorDeID(wg *sync.WaitGroup, idsChan chan<- string, closedChans chan<- int) {

	for i := 0; i < 100; i++ {
		id := uuid.New()
		//recibimos el valor del id hacia la variable idsChan con el indice de la iteracion y el id
		idsChan <- fmt.Sprintf("%d. %s", i+1, id.String())
	}
	close(idsChan)
	closedChans <- 1
	wg.Done()
}

func logIDS(wg *sync.WaitGroup, idsChan <-chan string, fakeIdsChan <-chan string, closedChans chan int) {
	//y aqui recibimos la variable idsChan hacia la variable id
	//	for id := range idsChan {
	//		fmt.Println(id)
	//	}

	closeCounter := 0

	for {
		select {
		case id, abierto := <-idsChan:
			if abierto {
				fmt.Println("ID", id)
			}

		case id, abierto := <-fakeIdsChan:
			if abierto {
				fmt.Println("FAKEID", id)
			}

		case count, abierto := <-closedChans:
			if abierto {
				closeCounter += count
			}
		}

		if closeCounter == 2 {
			close(closedChans)
			break
		}

	}
	wg.Done()

}
