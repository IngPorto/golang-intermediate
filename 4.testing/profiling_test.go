/**

PROFILING ayuda a mejorar el rendimiento de las pruebas

> go test -cpuprofile="cou.out"  .\4.testing\profiling_test.go

interpretar el profiling

> go tool pprof ./cou.out		<-- 
> > top      								<-- Desempeño en seg del test
> > list Fibonacci					<-- Análisis de desempeño de la función
> > web											<-- Exportar el análisis en html
> > pdf											<-- Exportar el análisis en .pdf



*/

package main

import "testing"


func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func TestFibonacci(t *testing.T) {
	tables  := []struct{
		n int
		expected int
	} {
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		if Fibonacci(item.n) != item.expected {
			t.Errorf("Expected %d but got %d", item.expected, Fibonacci(item.n))
		}
	}
}