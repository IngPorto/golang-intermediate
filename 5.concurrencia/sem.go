// Buffered channels como semáforos

// Permiten limitar la cantidad de Goroutines que se puedan
// ejecutar al mismo tiempo

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func main(){
	c := make(chan int, 2) // tiene solo dos espacio, se va a bloquear cuando complete los espacios y en 2 segundo se desbloqua porque se libera con <- c
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething2(i, &wg, c)
	}

	// Dejo el programa en un estado de bloqueo hasta que llegue a cero
	wg.Wait()
}
*/

func doSomething2(i int, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Done %d\n", i)
	<-c // libero el canal
}
