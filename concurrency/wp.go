package main

import "fmt"

func main() {
	tasks := []int{2, 3, 4, 5, 7, 10, 12, 40}
	nWorkers := 3
	// donde se van a leer los resultados
	jobs := make(chan int, len(tasks))
	// donde se van a transmitir los resultados
	results := make(chan int, len(tasks))
	// asignamos los workers
	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}
	// es aqui cuando ponemos a trabajar a los workers
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
