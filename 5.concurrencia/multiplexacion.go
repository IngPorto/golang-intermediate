package main

import (
	"fmt"
	"time"
)

func main() {
	channel1 := make(chan int)
	channel2 := make(chan int)
	
	duration1 := 4 * time.Second
	duration2 := 2 * time.Second

	go doSomething3(duration1, channel1, 1)
	go doSomething3(duration2, channel2, 2)

	/*
	fmt.Println(<-channel1) // Siempre se imprime de primero porque el programa se bloquea
	// esperando a que retorne lo que hay en el canal 1
	fmt.Println(<-channel2)
	*/

	for i := 0; i < 2; i++ {
		select { // evito bloquear el programa, pero me facilita estar a la espera de ambos canales. For == 2
			case channelMsg1 := <-channel1:
				fmt.Println(channelMsg1)
			case channelMsg2 := <-channel2:
				fmt.Println(channelMsg2)
		}
	}
}

func doSomething3(i time.Duration, c chan<- int, param int) {
	time.Sleep(i)
	c <- param
}