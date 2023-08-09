package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Enployee struct {
	id int
}

type TemporaryEmployee struct {
	Person
	Enployee
	taxRate int
}

// Función creada para el Struct, la cual estará relacionada a 
// una INTERFAZ
func (tEmployee TemporaryEmployee) getMessage() string {
	return fmt.Sprintf("This is a Temporary Employee")
}

// HERENCIA
// Composición sobre la Herencia
type FullTimeEmployee struct {
	Person   // Campo anónimo, solo el tipo, sin propiedades
	Enployee //
}

// Función creadad para el Struct, la cual estará relacionada a 
// una INTERFAZ
func (ftEmployee FullTimeEmployee) getMessage() string {
	return fmt.Sprintf("This is a Fulltime Employee")
}

// En Go las interfaces no se implementan de manera explicita 
type PrintInfo interface {
	getMessage() string
}

// Esta es la función que se va a llamar como la implementación de la
// INTERFACE
func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func GetMessage__without_interface(p Person) {
	fmt.Printf("Hi %s with age %d \n", p.name, p.age)
}

func main() {
	// Herencia
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Néstor" // Propiedad de Person
	ftEmployee.age = 32        // Propiedad de Person
	ftEmployee.id = 00121      // Propiedad de employee

	fmt.Printf("%+v\n",ftEmployee)

	// Interface
	GetMessage__without_interface(ftEmployee.Person)
	
	tEmployee := TemporaryEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)
}