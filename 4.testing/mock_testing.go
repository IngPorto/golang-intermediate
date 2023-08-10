package main

import (
	"time"
)

type Person struct {
	Name string
	Age  int
	Dni string
}

type Employee struct {
	id int
	position string
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}
type FullTimeEmployee struct {
	Person   // Campo anónimo, solo el tipo, sin propiedades
	Employee //
}

var GetPersonByDNI = func (dni string) (*Person, error) {
	time.Sleep(2 * time.Second)
	// SELECT * FROM Persona WHERE DNI ....
	var p Person
	return &p, nil
}

var GetEmployeeById = func (id int) (*Employee, error) {
	time.Sleep(2 * time.Second)
	// SELECT * FROM Empleado WHERE ID ....
	var e Employee
	return &e, nil
}
func GetFullTimeEmployeeById(id int, dni string) (*FullTimeEmployee, error) {
	var ftEmployee = FullTimeEmployee{}

	employee, err := GetEmployeeById(id)
	if err != nil {
		return &ftEmployee, err
	}
	ftEmployee.Employee = *employee

	person, err2 := GetPersonByDNI(dni)
	if err2 != nil {
		return &ftEmployee, err2
	}
	ftEmployee.Person = *person

	return &ftEmployee, nil
}