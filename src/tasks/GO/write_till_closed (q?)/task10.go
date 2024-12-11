package main

import (
	"context"
	"fmt"
	"reflect"
	"time"
)

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	// code here
	out := make(chan interface{})

	go func() {
		defer close(out)
		// can be done without mark, because return will terminate only goroutine
		// LOOP:
		for {
			select {
			case <-ctx.Done():
				// break LOOP
				return

			case val, ok := <-in:
				if !ok {
					// break LOOP
					return
				}

				// TODO: If no one is reading form this chan - should here be a nested select or not?
				out <- val
			}
		}
	}()

	return out
}

func main() {
	mainContext, _ := context.WithTimeout(context.Background(), 100*time.Millisecond)
	ch := make(chan interface{})

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	var res []interface{}
	for v := range orDone(mainContext, ch) {
		res = append(res, v)
	}
	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	} else {
		fmt.Println("Program worker correctly")
	}
}
