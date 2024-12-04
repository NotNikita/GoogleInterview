package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 5)
	ch3 := make(chan int, 5)

	go func() {
		defer close(ch1)
		for i := range 5 {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	go func() {
		defer close(ch2)
		for i := range 5 {
			ch2 <- i + 5
			time.Sleep(150 * time.Millisecond)
		}
	}()
	go func() {
		defer close(ch3)
		for i := range 5 {
			ch3 <- i + 10
			time.Sleep(250 * time.Millisecond)
		}
	}()

	result := mergeChannels(wg, ch1, ch2, ch3)

	go func() {
		for item := range result {
			fmt.Println(item)
		}
	}()

	wg.Wait()
	close(result)
	fmt.Println("Main finished execution")
}

func mergeChannels(wg *sync.WaitGroup, chs ...<-chan int) chan int {
	result := make(chan int, 15)

	// for each chan we create dedicated worker to process it data
	doWork := func(ch <-chan int) {
		defer wg.Done()
		for it := range ch {
			result <- it
		}
	}

	wg.Add(len(chs))
	for _, channel := range chs {
		go doWork(channel)
	}
	return result
}
