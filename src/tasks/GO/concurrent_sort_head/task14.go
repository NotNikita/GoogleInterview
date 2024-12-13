package main

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
	"sync"
)

func main() {
	f1 := `aaa
ddd `
	f2 := `bbb eee
`
	f3 := `ccc
fff `

	files := []io.Reader{
		strings.NewReader(f1),
		strings.NewReader(f2),
		strings.NewReader(f3),
	}

	rows, err := ConcurrentSortHead(4, files...)
	if err != nil {
		panic(err)
	}
	if !reflect.DeepEqual(rows, []string{"aaa", "bbb", "ccc", "ddd"}) {
		fmt.Println("Results dont match", rows)
		panic("wrong code")
	} else {
		fmt.Println("All good, mate")
	}
}

type FileReadInfo struct {
	ReaderIndex int
	Result      string
}

// Should return "m" first strings. Reading should be concurrent
// files are readers, sorted in ascending order
func ConcurrentSortHead(m int, files ...io.Reader) ([]string, error) {
	// code here
	readingsCh := make(chan FileReadInfo, m)
	var err error
	wg := sync.WaitGroup{}

	wg.Add(len(files))
	for i := 0; i < len(files); i++ {

		go func() {
			defer wg.Done()

			scanner := bufio.NewScanner(files[i])
			scanner.Split(bufio.ScanWords)
			counter := 0
			for scanner.Scan() {
				readingsCh <- FileReadInfo{
					ReaderIndex: i + counter*len(files),
					Result:      scanner.Text(),
				}
				counter++
			}

		}()
	}

	go func() {
		wg.Wait()
		close(readingsCh)
	}()

	// Parse values from "readingsCh" in correct order
	result := make([]string, m)

	for info := range readingsCh {
		fmt.Println("read from chan", info)
		if info.ReaderIndex > 3 {
			continue
		}
		result[info.ReaderIndex] = info.Result
	}
	return result, err
}
