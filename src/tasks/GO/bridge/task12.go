package main

import (
	"context"
	"fmt"
	"reflect"
)

func main() {
	genVals := func() <-chan <-chan interface{} {
		out := make(chan (<-chan interface{}))
		go func() {
			defer close(out)
			for i := 0; i < 3; i++ {
				stream := make(chan interface{}, 1)
				stream <- i
				close(stream)
				out <- stream
			}
		}()
		return out
	}

	var res []interface{}
	for v := range bridge(context.Background(), genVals()) {
		res = append(res, v)
	}

	if !reflect.DeepEqual(res, []interface{}{0, 1, 2}) {
		panic("wrong code")
	} else {
		fmt.Println("Program worker correctly")
	}
}

/*
Consequetivly reads channels from "ins", and transfer data from them into single out channel

- Function is live, only while "ins" is open and context is not cancelled
- "Read channel" is parsed, only while it is open and context is not cancelled
*/
func bridge(ctx context.Context, ins <-chan <-chan interface{}) <-chan interface{} {
	// code here
	out := make(chan interface{}, 1)

	go func() {
		defer close(out)

		for ch := range ins {

			for val := range orDone(ctx, ch) {
				select {
				case <-ctx.Done():
					return

				case out <- val:
				}
			}

		}
	}()

	return out
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
				select {
				case out <- v:
				case <-ctx.Done():

				}
			}
		}
	}()
	return out
}
