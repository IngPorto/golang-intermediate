package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee *FullTimeEmployee
	}{
		{
			id:  1,
			dni: "12345678",
			mockFunc: func() {
				GetEmployeeById = func(id int) (*Employee, error) {
					return &Employee{
						id:       1,
						position: "Developer",
					}, nil
				}
				GetPersonByDNI = func(dni string) (*Person, error) {
					return &Person{
						Name: "Néstor",
						Age:  32,
						Dni:  "324",
					}, nil
				}
			},
			expectedEmployee: &FullTimeEmployee{
				Person: Person{
					Age:  32,				// Pasa
					Name: "Néstor",	// Pasa
					Dni:  "300",		// Falla, es 324
				},
				Employee: Employee{
					id: 1,
					position: "Developer",
				},
			},
		},
	}

	// Guardo las funciones originales que están en mock_testing.go
	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDni := GetPersonByDNI

	for _, test := range table {
		// Sobreescribo GetEmployeeById y GetPersonByDNI 
		test.mockFunc()
		ft, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err !=nil {
			t.Errorf("Error when getting Employee: %s", err)
		}

		if ft.Age != test.expectedEmployee.Age {
			t.Errorf("Expected %d but got %d", test.expectedEmployee.Age, ft.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDni

}