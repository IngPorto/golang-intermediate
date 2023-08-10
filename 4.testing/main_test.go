// TEST
// Es necesario que el proyecto esté inicializado como módulo
// go mod init

// CORRER EL TEST
// go test ./<ruta de mi proyecto>

// PAQUETE RECOMENDADO PARA CORRER TEST
// go get -u github.com/stretchr/testify

package main

import "testing"

func TestSum(t *testing.T){
	total := Sum(5, 5) // Sum está definido en el archivo main.go
	expected := 10

	if total != expected {
		t.Errorf("Expected %d but got %d", expected, total)
	}

}

func TestSumWithSlice(t *testing.T) {
	tables := []struct{ // https://github.com/a8m/golang-cheat-sheet#structs
		a int
		b int
		expected int
	} {
		{1,2,3},
		{4,5,9},
		{10,20,30},
	}

	for _, item := range tables {
		total := Sum(item.a, item.b)
		if total != item.expected {
			t.Errorf("Expected %d but got %d", item.expected, total)
		}
	}
}