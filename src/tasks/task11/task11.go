package main

import (
	"fmt"
	"time"
)

/*
Реализовать функцию выполняющую батчинг значений из канала с
Для возврата батчей использовать выходной канал  chan []any
c - входной канал значений. Из этих значений нужно сформировать батчи
batchSize - размер батча
*/
func doBatching(c chan int, batchSize int) chan []int {
	// Write code here
	out := make(chan []int, 10)
	var payload []int

	go func() {
		defer close(out)
		for val := range c {
			payload = append(payload, val)

			if len(payload) == batchSize {
				out <- payload
				payload = []int{}
			}
		}

		if len(payload) != 0 {
			out <- payload
		}
	}()

	return out
}

// batchTimeout - интервал отправки batch в out channel
func doBatchingWithtimeout(c chan int, batchSize int, batchTimeout time.Duration) chan []int {
	// Write code here
	out := make(chan []int, 10)
	ticker := time.NewTicker(batchTimeout)
	defer ticker.Stop()
	var payload []int

	go func() {
		defer close(out)

	LOOP:
		for {
			select {
			case val, ok := <-c:
				if !ok {
					break LOOP
				}
				payload = append(payload, val)

				if len(payload) == batchSize {
					out <- payload
					payload = []int{}
					ticker.Reset(batchTimeout)
				}
			case <-ticker.C:
				if len(payload) != 0 {
					out <- payload
					payload = []int{}
				}
			}

		}
		if len(payload) != 0 {
			out <- payload
		}
	}()

	return out
}

func main() {
	c := make(chan int, 10)
	batchSize := 2
	batchTimeout := 2 * time.Second

	// c <- 1
	// c <- 2
	// c <- 3
	// c <- 4
	// c <- 5
	// close(c)
	// out := doBatching(c, 2)
	out := doBatchingWithtimeout(c, batchSize, batchTimeout)

	go func() {
		for i := 1; i <= 5; i++ {
			c <- i
			time.Sleep(500 * time.Millisecond) // Simulate some delay between sends
		}
		close(c)
	}()

	go func() {
		for value := range out {
			fmt.Println("received at", time.Now().Format(time.TimeOnly), value)
		}
	}()

	time.Sleep(time.Minute)
}
