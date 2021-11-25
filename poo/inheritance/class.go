package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full time Employee"
}

type TemporaryEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporaryEmployee) getMessage() string {
	return "Temporary time Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan"
	ftEmployee.age = 29
	ftEmployee.id = 1
	fmt.Println(ftEmployee)
	tEmployee := TemporaryEmployee{}
	tEmployee.name = "Celina"
	tEmployee.age = 66
	tEmployee.id = 2
	fmt.Println(tEmployee)
	fmt.Println("sin interfaz=", tEmployee.getMessage())
	getMessage(ftEmployee)
	getMessage(tEmployee)
	//GetMessage(ftEmployee) // no es posible, porque es composicion sobre herencia
}
