package main

import (
	"fmt"
	gr "intermediate/1.repaso/go_routines"
	"strconv"
)

func main() {
	var x int
	x = 8
	y := 7
	fmt.Println(x)
	fmt.Println(y)

	// No hay excepciones, no hay try y catch. Hay que manejar
	// los errores de forma explícita.
	valor, err := strconv.ParseInt("8", 10, 8)
	if err != nil {
		fmt.Printf("Error: %v\n", err) 
	} else {
		fmt.Println(valor)
	}

	// Mapas
	m := make(map[string]int)
	m["edad"] = 32
	fmt.Println(m["edad"])

	// Slice
	s := []int{1, 2, 3, 4, 5}
	for index, value := range s {
		fmt.Println(index, value)
	}
	s = append(s, 6)

	// Go-routines
	// canal para poder comunicar las diferentes rutinas
	c := make(chan int)
	go gr.DoSomething(c)
	// Esperar a que la gorutina termine
	<- c

	// Apuntadores
	g:= 25	
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}