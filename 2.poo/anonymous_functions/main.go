package main

import (
	"fmt"
	"time"
)

func main() {
	// Tipo 1
	x := 5

	y := func() int {
		return x * 2
	}()

	fmt.Println(y)

	// Tipo 2
	c := make(chan int)
	go func() {
		fmt.Println("Starting the function")
		time.Sleep(5 * time.Second)
		fmt.Println("Ending the function")
		c <- 1
	}()
	<- c
}