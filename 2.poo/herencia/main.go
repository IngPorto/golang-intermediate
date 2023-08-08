package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Enployee struct {
	id int
}

// Composición sobre la Herencia
type FullTimeEmployee struct {
	Person   // Campo anónimo, solo el tipo, sin propiedades
	Enployee //
}

func GetMessage(p Person) {
	fmt.Printf("Hi %s with age %d \n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Néstor" // Propiedad de Person
	ftEmployee.age = 32        // Propiedad de Person
	ftEmployee.id = 00121      // Propiedad de employee

	fmt.Printf("%+v\n",ftEmployee)
	GetMessage(ftEmployee.Person)
}