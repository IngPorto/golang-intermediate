package main

import (
	"fmt"
	e "intermediate/2.poo/employee"
)

func constructores() {
	// Instanciación 2
	emp := e.Employee{
		Id:   00121,
		Name: "Néstor",
		Vacation: true,
	}
	fmt.Printf("%+v\n", emp)

	// Instanciación 3
	emp3 := new(e.Employee)
	fmt.Printf("%+v\n", emp3) // devuelve un &, una referencia a la memoria
	fmt.Printf("%+v\n", *emp3) // devuelve el valor

}

// Instanciación 4
// Es muy paracido al New, porque devuelve una referencia en memoria (&)
// Así se puede modificar el valor de esa referencia en memoria sin
// generar más información en memoria. 

// NOTA: Cuando trabajamos con structs y los pasamos a funciones Go los
// trata como copias, así que hay que devolver la referencias
func NewEmployee(id int, name string, vacation bool) *e.Employee {
	return &e.Employee{
		Id:   id,
		Name: name,
		Vacation: vacation,
	}
	
}

func main() {
	// Instanciación 1
	emp1 := e.Employee{}
	fmt.Printf("%+v\n", emp1)

	emp1.Id = 12321
	emp1.Name = "Juan"
	fmt.Printf("%+v\n", emp1)

	// uso del receiver function
	// el struct se envía a sí mismo como parámetro, así el método 
	// SetId puede acceder a las propiedades de Employee
	emp1.SetId(44446)
	emp1.SetName("Pedro")
	fmt.Printf("%+v\n", emp1)

	constructores()
	emp4 := NewEmployee(555656, "Sergio", true)
	fmt.Printf("%+v\n", *emp4)
}