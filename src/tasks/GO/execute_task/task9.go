package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const timeout = 100 * time.Millisecond

func main() {
	ctx, _ := context.WithTimeout(context.Background(), timeout)
	err := executeTaskWithTimeout(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("task done")
}

func executeTaskWithTimeout(ctx context.Context) error {
	// code here
	taskFinished := make(chan struct{})

	go func() {
		defer close(taskFinished)
		executeTask()
		taskFinished <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-taskFinished:
		return nil
	}
}

func executeTask() {
	time.Sleep(time.Duration(rand.Intn(3)) * timeout)
}
