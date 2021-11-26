package main

import "fmt"

// canal solo de entrada, se especifica <- a la derecha
func Generator(c chan<- int) {
	for i := 1; i <= 10; i++ {
		c <- i
	}
	close(c) //cierra el canal
}

// in es de salida <-izquierda, out es de entrada derecha<-
func Double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- 2 * value
	}
	close(out)
}

// c es de salida <-izquierda
func Print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generator := make(chan int)
	doubles := make(chan int)
	go Generator(generator)
	go Double(generator, doubles)
	Print(doubles)
}
