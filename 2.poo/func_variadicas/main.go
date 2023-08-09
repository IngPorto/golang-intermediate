/**
 Funciones variádicas y retornos con nombre
*/
package main

import "fmt"

func sum(a, b int) int {
	return a + b
}

func sum_variadica(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func printNames(names ...string) {
	for _, name := range names {
		fmt.Println(name)
	}
}

// Retornos con nombres
func getValues(x int) (elDoble int, elTriple int, elCuadruple int) {
	elDoble = x * 2
	elTriple = x * 3
	elCuadruple = x * 4
	return
}

func main() {
	fmt.Println(sum(1, 2))
	printNames("Juan", "Pedro", "Maria")
	fmt.Println(sum_variadica(1, 2, 3, 4, 5))

	// Torno con nombres
	fmt.Println(getValues(2))
}