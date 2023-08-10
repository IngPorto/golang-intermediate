package main

import "fmt"

func main(){
	tasks := []int{ 2, 4, 7, 10, 12, 32, 40}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++{
		go Worker(i, jobs, results)
	}

	for _, task := range tasks{
		jobs <- task
	}
	close(jobs)

	for a := 0; a < len(tasks); a++ {
		<-results
	}
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d starting job %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker %d finished job %d\n and fib is %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
