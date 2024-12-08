package main

import (
	"context"
	"fmt"
	"math/rand"
)

// infinitly calling provided function "fn" and putting it result into chan
func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{}, 10)
	// code here

	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return

			case out <- fn():
			}
		}
	}()

	return out
}

// Reads N values from in channel, where N <= num, and put read values into output chan.
// Works while in is open or context is not cancelled
func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	// code here
	out := make(chan interface{}, num)
	count := 0

	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-ctx.Done():
				return
			case out <- i:
				count++
				if count == num {
					return
				}
			}
		}
	}()

	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rand := func() interface{} { return rand.Int() }
	var res []interface{}
	for num := range take(ctx, repeatFn(ctx, rand), 13) {
		res = append(res, num)
	}
	if len(res) != 13 {
		panic("wrong code")
	}
	fmt.Println("Results obtained:", res)

	fmt.Scan()
}
