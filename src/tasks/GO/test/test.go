package main

import "fmt"

func main() {
	arr := make([]int, 3, 4)
	sm := arr[:1]

	appendSlice(sm)
	fmt.Println("result", arr)
}

func appendSlice(slice []int) {
	slice = append(slice, 1)
}
