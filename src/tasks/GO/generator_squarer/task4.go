package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	pipeline := squarer(ctx, generator(ctx, 1, 2, 3))
	for x := range pipeline {
		fmt.Println(x)
	}
}

// Should write spreaded slice of ints into channel
func generator(ctx context.Context, in ...int) <-chan int {
	ch := make(chan int, len(in))
	// напишите ваш код здесь

	go func() {
		defer close(ch)
	LOOP:
		for _, v := range in {
			select {
			case <-ctx.Done():
				break LOOP
			case ch <- v:
			}
		}
	}()

	return ch
}

// Should receive chan of ints, after it reads each
// element and writes square of it into output channel
func squarer(ctx context.Context, in <-chan int) <-chan int {
	// напишите ваш код здесь
	out := make(chan int, 5)

	go func() {
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case v, ok := <-in:
				if !ok {
					close(out)
					break LOOP
				}
				out <- v * v
			}

		}
	}()

	return out
}
