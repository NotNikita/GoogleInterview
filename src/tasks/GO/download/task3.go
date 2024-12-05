/*
Download func
- input: urls for downloading
- concurrently fetches info from each url (use fakeDownload for downloading)
- if fakeDownload returns errors - we need to return all of them (errors.Join)
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// timeoutLimit - вероятность, с которой не будет возвращаться ошибка от fakeDownload():
// timeoutLimit = 100 - ошибок не будет;
// timeoutLImit = 0
// Можете "поиграть"  - всегда будет возвращаться ошибка.
// с этим параметром, для проверки случаев с возвращением ошибки.
const timeoutLimit = 80

type Result struct {
	msg string
	err error
}

// fakeDownload - имитирует разное время скачивания для разных адресов
func fakeDownload(url string) Result {
	r := rand.Intn(100)
	time.Sleep(time.Duration(r) * time.Millisecond)
	if r > timeoutLimit {
		return Result{
			err: errors.New(fmt.Sprintf("failed to download data from %s: timeout", url)),
		}
	}
	return Result{
		msg: fmt.Sprintf("downloaded data from %s\n", url),
	}
}

// downloadWg - utilize channel and wait group
func downloadWg(urls []string) ([]string, error) {
	// code here
	n := len(urls)
	var err error
	wg := &sync.WaitGroup{}
	results := make(chan Result, n)
	returns := make([]string, 0, n)

	wg.Add(n)
	for _, url := range urls {
		go func() {
			defer wg.Done()
			results <- fakeDownload(url)

		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for msg := range results {
		if msg.err != nil {
			err = errors.Join(err, msg.err)
		}
		returns = append(returns, msg.msg)
	}

	return returns, err
}

func downloadCtx(ctx context.Context, urls []string) ([]string, error) {
	// code here
	n := len(urls)
	var err error

	wg := &sync.WaitGroup{}
	results := make(chan Result, n)
	returns := make([]string, 0, n)

	wg.Add(n)
	for _, url := range urls {
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			case results <- fakeDownload(url):
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for msg := range results {
		if msg.err != nil {
			err = errors.Join(err, msg.err)
		}
		returns = append(returns, msg.msg)
	}

	return returns, err
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	urls := []string{
		"https://example.com/e25e26d3-6aa3-4d79-9ab4-fc9b71103a8c.xml", "https://example.com/a601590e-31c1-424a-8ccc-decf5b35c0f6.xml", "https://example.com/1cf0dd69-a3e5-4682-84e3-dfe22ca771f4.xml", "https://example.com/ceb566f2-a234-4cb8-9466-4a26f1363aa8.xml", "https://example.com/b6ed16d7-cb3d-4cba-b81a-01a789d3a914.xml",
	}

	// 1. Download Wait Group
	// msgs, err := downloadWg(urls)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(msgs)

	// 2. Download Context with Cancel
	msgs, err := downloadCtx(ctx, urls)
	if err != nil {
		fmt.Println("downloadCtx error:", err)
	}
	fmt.Println("downloadCtx results:")
	fmt.Println(msgs)
}
