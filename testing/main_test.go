package main

import "testing"

// Para correr el test: go test, pero primero
// se debe inicializar como módulo este directorio

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
