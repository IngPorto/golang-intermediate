package main

import (
	"fmt"
	e "intermediate/2.poo/employee"
)



func main() {
	emp1 := e.Employee{}
	fmt.Printf("%+v", emp1)

	emp1.Id = 12321
	emp1.Name = "Juan"
	fmt.Printf("%+v", emp1)

	// uso del receiver function
	// el struct se envía a sí mismo como parámetro, así el método 
	// SetId puede acceder a las propiedades de Employee
	emp1.SetId(44446)
	emp1.SetName("Pedro")
	fmt.Printf("%+v", emp1)
}