// Definiendo channels de lectura y escritura

package main

func Generator (c chan<- int) {
	for i := 1; i < 10; i++ {
		c <- i
	}
	close(c)
}

func Double (in <-chan int, out chan int ) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

func Print (c chan int){
	for value := range c {
		println(value)
	}
}

func ____main (){
	generator := make(chan int)
	double := make(chan int)

	go Generator(generator)
	go Double(generator, double)
	Print(double) // el programa se bloquea en esta linea
}