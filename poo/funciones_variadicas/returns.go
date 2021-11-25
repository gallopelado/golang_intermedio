package main

import "fmt"

// al usar este tipo de entrada se considerará como un slice
// para cualquier tipo de dato
func sumar(values ...int) int {
	var resultado int
	for _, value := range values {
		resultado = resultado + value
	}
	return resultado
}

// podemos escoger el nombre de los retornos
// go va inferir automáticamente lo que vamos a enviar
func getValues(x int) (double, triple, quad int) {
	double = 2 * x
	triple = 3 * x
	quad = 4 * x
	return
}

func main() {
	fmt.Println(sumar(1, 10, 5))
	fmt.Println(getValues(2))
}
