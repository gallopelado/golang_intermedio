package main

import "testing"

// Para correr el test: go test, pero primero
// se debe inicializar como módulo este directorio

/*
	Tratar el coverage
	go test -cover
	PASS
	coverage: 25.0% of statements
	ok      github.com/gallopelado/testing  0.001s

	En este caso no realizamos el test a otras funciones

	Crea un archivo de funciones
	go test -coverprofile=coverage.out

	Verificar el archivo
	go tool cover -func=coverage.out
	github.com/gallopelado/testing/main.go:3:       Sum             100.0%
	github.com/gallopelado/testing/main.go:7:       GetMax          0.0%
	total:                                          (statements)    25.0%

	Si lo queremos ver más bonito en un browser
	go tool cover -html=coverage.out
*/

func TestSum(t *testing.T) {
	// total := Sum(5, 5)
	// if total != 10 {
	// 	t.Errorf("Sum was incorrect, got %d expected %d", total, 10)
	// }
	tables := []struct {
		a int
		b int
		n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 26, 51},
	}
	for _, value := range tables {
		total := Sum(value.a, value.b)
		if total != value.n {
			t.Errorf("Sum was incorrect, got %d expected %d", total, value.n)
		}
	}
}

func TestMax(t *testing.T) {
	tables := []struct {
		a int
		b int
		n int
	}{
		{4, 2, 4},
		{3, 2, 3},
		{72, 100, 100},
	}
	for _, item := range tables {
		max := GetMax(item.a, item.b)
		if max != item.n {
			t.Errorf("GetMax was incorrect, got %d expected %d", max, item.n)
		}
	}
}

/*
	Para BigO notation
	Ver el rendimiento
	go test -cpuprofile=cpu.out

	Revisamos el fichero generado
	go tool pprof cpu.out
	Comandos:
	top
	list Fibonacci
	web(Instalar sudo apt install graphviz)
*/

func TestFib(t *testing.T) {
	tables := []struct {
		a int
		n int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}
	for _, item := range tables {
		fib := Fibonacci(item.a)
		if fib != item.n {
			t.Errorf("Fibonacci was incorrect, got %d expected %d", fib, item.n)
		}
	}
}
