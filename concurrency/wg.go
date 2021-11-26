package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) //contador del wg, incrementa en 1
		// enviado como puntero o direccion de memoria, referenciado
		go doSomething(i, &wg)
	}
	wg.Wait() //bloquea hasta que el contador este en cero
}

//recibido como parametro
func doSomething(i int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementa en uno
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Println("End")
}
