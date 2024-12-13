package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

const (
	MAX_PARALLEL_FUNCS = 2
)

type waiter interface {
	wait() error
	run(ctx context.Context, f func(ctx context.Context) error)
}

// "waitGroup" should implement "waiter" interface
type waitGroup struct {
	// code here
	workerWg sync.WaitGroup
	mu       sync.Mutex

	workers chan struct{}
	errors  []error
}

// Should return an error, if one of "run" functions finished with error
// If there multiple instances of such errors, return "errors.Join"
func (g *waitGroup) wait() error {
	// code here
	var err error

	g.workerWg.Wait()
	close(g.workers)
	for _, er := range g.errors {
		if er != nil {
			err = errors.Join(err, er)
		}
	}

	return err
}

// Should concurrently run functions inside "run", passed via "f"
func (g *waitGroup) run(ctx context.Context, fn func(ctx context.Context) error) {
	// code here
	g.workers <- struct{}{}
	g.workerWg.Add(1)

	go func() {
		defer func() {
			<-g.workers
			g.workerWg.Done()
		}()

		if maybeErr := fn(ctx); maybeErr != nil {
			g.mu.Lock()
			g.errors = append(g.errors, maybeErr)
			g.mu.Unlock()
		}
	}()
}

// Struct constructor
// Should create create N instances of "waitGroup", where N is <= maxParallel
func NewGroupWait(maxParallel int) waiter {
	// code here
	return &waitGroup{
		workerWg: sync.WaitGroup{},
		mu:       sync.Mutex{},
		workers:  make(chan struct{}, maxParallel),
		errors:   []error{},
	}
}

func main() {
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
}
