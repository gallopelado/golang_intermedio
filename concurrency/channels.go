package main

import "fmt"

func main() {
	// Canales sin buffer, se bloqueara
	//c := make(chan int)
	// Canales con buffer, no se bloqueara
	c := make(chan int, 3)
	// enviamos un valor a ese canal
	c <- 1
	c <- 2
	c <- 3
	c <- 4 //este salio del limite
	// Se producirá un error en la siguiente línea
	fmt.Println(<-c)
}
