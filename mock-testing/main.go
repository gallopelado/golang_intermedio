package main

import "time"

type Person struct {
	DNI  string
	Name string
	Age  int
}

type Employee struct {
	Id       int
	Position string
}

type FullTimeEmployee struct {
	Employee
	Person
}

// para no usar func, al principio ponemos var

var GetPersonByDNI = func(dni string) (Person, error) {
	time.Sleep(5 * time.Second)
	return Person{}, nil
}

var GetEmployeeById = func(id int) (Employee, error) {
	time.Sleep(5 * time.Second)
	return Employee{}, nil
}

func GetFullTimeEmployeeById(id int, dni string) (FullTimeEmployee, error) {
	var ftEmployee FullTimeEmployee
	// Traer empleado
	e, err := GetEmployeeById(id)
	if err != nil {
		return ftEmployee, err
	}
	ftEmployee.Employee = e
	// Traer persona
	p, perr := GetPersonByDNI(dni)
	if perr != nil {
		return ftEmployee, err
	}
	ftEmployee.Person = p
	// No hubieron errores, entonces retornamos toda la estructura
	return ftEmployee, nil
}
