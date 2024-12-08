package main

import (
	"fmt"
	"sync"
)

const (
	numJobs    = 5
	numWorkers = 3
)

type JobResult struct {
	DoneByWorker int
	Result       int
}

func worker(wg *sync.WaitGroup, id int, f func(int) int, jobs <-chan int, results chan<- JobResult) {
	// code here
	defer wg.Done()
	for job := range jobs {
		jobRes := f(job)
		results <- JobResult{
			DoneByWorker: id,
			Result:       jobRes,
		}
	}
}

// write a Worker pool
// You need to perform "numJobs" concurrently, using "numWorkers" workers
// Each worker should be started only once, in main func
// Each worker takes job from jobs chan, and writes result of "f(job)" into results chan

func main() {
	jobs := make(chan int, numJobs)
	results := make(chan JobResult, numJobs)
	wg := sync.WaitGroup{}
	multiplier := func(x int) int {
		return x * 10
	}
	// code here

	// Starting workers
	wg.Add(numWorkers)
	for workerId := range numWorkers {
		go worker(&wg, workerId, multiplier, jobs, results)
	}

	// populate jobs
	go func() {
		defer close(jobs)
		for i := range numJobs {
			jobs <- i
		}
	}()

	doneSignal := make(chan struct{})

	// async read from results
	go func() {
		defer close(doneSignal)
		for res := range results {
			fmt.Printf("Result of %v, done by %v \n", res.Result, res.DoneByWorker)
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	<-doneSignal
}

// Explanation
/*
1. Populating jobs
2. Closing jobs
3. After jobs is closed - workers are terminated, as they cant read from closed chan
4. wg.Wait() fired, and results closed
5. As results closed - doneSignal fired
6. Main finished
*/
