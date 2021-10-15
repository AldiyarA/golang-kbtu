package main

import (
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id    int
	jobs  chan int
	power int
}

func newWorker(id int, jobs chan int, power int) *Worker {
	return &Worker{id, jobs, power}
}
func (worker *Worker) Work() {
	for job := range worker.jobs {
		fmt.Printf("Worker #%d started work %d\n", worker.id, job)
		x := job
		for x > 0 {
			time.Sleep(time.Second)
			x -= worker.power
		}
		fmt.Printf("Worker #%d finished work %d\n", worker.id, job)
	}
}

func main() {
	jobs := make(chan int)
	workersNumber := 4
	jobsNumber := 20
	wg := new(sync.WaitGroup)
	wg.Add(workersNumber)
	for i := 1; i <= workersNumber; i++ {
		go func(i int) {
			newWorker(i, jobs, i).Work()
			wg.Done()
		}(i)
	}
	for i := 1; i <= jobsNumber; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
