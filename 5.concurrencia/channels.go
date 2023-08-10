package main

import "fmt"

/*
func main() {
	canalSinBuffer()
}
*/

func canalSinBuffer() {
	c := make(chan int)

	c <- 1 // Se bloquea en esta línea
	// fatal error: all goroutines are asleep - deadlock!

	fmt.Println(<-c) // Espera retornar la información pero está bloqueado anteriomente
}

func canalConBuffer() {
	c := make(chan int, 1)

	c <- 1 // No se bloquea 

	fmt.Println(<-c) // Retorna
}