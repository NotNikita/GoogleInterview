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
func doBatching(c chan int, batchSize int) chan []int{
	// Write code here
	out := make(chan []int, 10)
	var payload []int

	go func(){
		defer close(out)
		for val := range c {
			payload = append(payload, val)
	
			if (len(payload) == batchSize) {
				out <- payload
				payload = []int{}
			}
		}
	
		if (len(payload) != 0) {
			out <- payload
		}
	}()
	
	

	return out
}

func main() {
	c := make(chan int, 10)
	c <- 1
	c <- 2
	c <- 3
	c <- 4
	c <- 5
	close(c)

	out := doBatching(c, 2)

	go func() {
		for value := range out {
			fmt.Println(value)
		}
	}()

	time.Sleep(time.Second)
}