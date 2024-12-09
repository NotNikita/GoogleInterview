package main

import (
	"fmt"
)

// Return sorted chanC, which contains sorted elements from both channels
func mergeSorted(a, b <-chan int) <-chan int {
	// code here
	chanC := make(chan int)

	go func() {
		defer close(chanC)

		av, ok1 := <-a
		bv, ok2 := <-b

		for ok1 && ok2 {
			if av > bv {
				chanC <- bv
				bv, ok2 = <-b
			} else {
				chanC <- av
				av, ok1 = <-a
			}
		}

		for ok1 {
			chanC <- av
			av, ok1 = <-a
		}
		for ok2 {
			chanC <- bv
			bv, ok2 = <-b
		}

	}()

	return chanC
}

// sorted chanA
func fillChanA(c chan int) {
	c <- 1
	c <- 2
	c <- 4
	close(c)
}

// sorted chanB
func fillChanB(c chan int) {
	c <- -1
	c <- 4
	c <- 5
	close(c)
}

func main() {
	a, b := make(chan int), make(chan int)
	go fillChanA(a)
	go fillChanB(b)
	c := mergeSorted(a, b)
	for val := range c {
		fmt.Println("value", val)
	}
}
