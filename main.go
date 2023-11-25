package main

import (
	"log"
	"time"
)

const numJobs = 10

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		log.Println("Result:", <-results)
	}

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		log.Printf("Worker %d processing job %d\n", id, j)
		time.Sleep(time.Second)
		results <- j
	}
}
