package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	i := 0
	inc := func() interface{} {
		i++
		return i
	}
	out1, out2 := tee(ctx, take(ctx, repeatFn(ctx, inc), 3))
	var res1, res2 []interface{}

	for val1 := range out1 {
		res1 = append(res1, val1)
		res2 = append(res2, <-out2)
	}
	exp := []interface{}{1, 2, 3}
	if !reflect.DeepEqual(res1, exp) || !reflect.DeepEqual(res2, exp) {
		panic("wrong code")
	} else {
		fmt.Println("Program worker correctly")
	}
}

// Incoming data from channel in is sent to two channels, that are returned here
// use orDone func from "write_till_closed" task for simplicity
func tee(ctx context.Context, in <-chan interface{}) (_, _ <-chan interface{}) {
	// code here
	out1 := make(chan interface{}, 1)
	out2 := make(chan interface{}, 1)

	go func() {
		defer close(out2)
		defer close(out1)

		for val := range orDone(ctx, in) {
			var a, b = out1, out2
			// 2 here is amount of channels (we return 2, so its two)
			for i := 0; i < 2; i++ {
				select {
				case <-ctx.Done():
					return
				// Because we dont guarantee what happens first, we nullify the channel
				case a <- val:
					a = nil
				case b <- val:
					b = nil
				}
			}
		}

		// for {
		// 	select {
		// 	case <-ctx.Done():
		// 		return
		// 	case val, ok := <-in:
		// 		if !ok {
		// 			return
		// 		}

		// 		ch1 <- val
		// 		ch2 <- val
		// 	}
		// }
	}()

	return out1, out2
}

func orDone(ctx context.Context, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				// TODO: no dounle-nested select works too! ±± But not safe?
				select {
				case out <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return out
}

func repeatFn(ctx context.Context, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
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

func take(ctx context.Context, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				out <- v
			}
		}
	}()
	return out
}
