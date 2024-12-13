package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	MAX_PARALLEL_FUNCS = 2
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

// v2: Worker Pool
// "waitGroup" should implement "waiter" interface
type waitGroup struct {
	// code here
	mu        sync.Mutex
	workersWg sync.WaitGroup
	tasksWg   sync.WaitGroup

	taskQueue chan task
	errors    []error
}

type task struct {
	ctx context.Context
	fn  func(context.Context) error
}

// Should return an error, if one of "run" functions finished with error
// If there multiple instances of such errors, return "errors.Join"
func (g *waitGroup) wait() error {
	// code here
	g.tasksWg.Wait()
	close(g.taskQueue)
	g.workersWg.Wait()

	var outErr error

	for _, er := range g.errors {
		outErr = errors.Join(outErr, er)
	}

	return outErr
}

// v2: Add task to task queue
// Should concurrently run functions inside "run", passed via "f"
func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	// code here

	g.tasksWg.Add(1)
	g.taskQueue <- task{
		ctx: ctx,
		fn:  fn,
	}

}

func worker(theWaitGroup *waitGroup) {
	defer theWaitGroup.workersWg.Done()

	for task := range theWaitGroup.taskQueue {
		if maybeError := task.fn(task.ctx); maybeError != nil {
			theWaitGroup.mu.Lock()
			theWaitGroup.errors = append(theWaitGroup.errors, maybeError)
			theWaitGroup.mu.Unlock()
		}
		theWaitGroup.tasksWg.Done()
	}
}

// v2: create Worker pool with "maxParallel" workers
// Struct constructor
// Should create create N instances of "waitGroup", where N is <= maxParallel
func NewGroupWait(maxParallel int) waiter {
	// code here
	theWaitGroup := &waitGroup{
		mu:        sync.Mutex{},
		workersWg: sync.WaitGroup{},
		tasksWg:   sync.WaitGroup{},
		taskQueue: make(chan task),
		errors:    []error{},
	}

	theWaitGroup.workersWg.Add(maxParallel)
	for i := 0; i < maxParallel; i++ {
		go worker(theWaitGroup)
	}

	return theWaitGroup
}

func main() {
	start := time.Now()

	ctx := context.Background()

	g := NewGroupWait(MAX_PARALLEL_FUNCS)
	expErr1 := errors.New("got error 1")
	expErr2 := errors.New("got error 2")

	g.run(ctx, func(ctx context.Context) error {
		return nil
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr2
	})
	g.run(ctx, func(ctx context.Context) error {
		return expErr1
	})

	err := g.wait()
	if !errors.Is(err, expErr1) || !errors.Is(err, expErr2) {
		fmt.Println("IF: Our errors:", err)
		panic("wrong code")
	} else {
		fmt.Println("ELSE: Our errors:", err)
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed) // 1,95s -> 3,6s
}
