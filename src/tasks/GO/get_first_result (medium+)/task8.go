package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type result struct {
	msg string
	err error
}
type search func() *result
type replicas []search

func fakeSearch(kind string) search {
	return func() *result {
		time.Sleep(time.Duration(rand.Intn(130)) * time.Millisecond)
		return &result{
			msg: fmt.Sprintf("%q result", kind),
		}
	}
}

func worker(ctx context.Context, wg *sync.WaitGroup, foo search, out chan<- *result) {
	defer wg.Done()

	select {
	case <-ctx.Done():
		return
	case out <- foo():
		return
	}
}

/*
runs a concurrent search in replicas, that returns first available result from it
return ctx error, if ctx finished earlier
*/
func getFirstResult(ctx context.Context, replicas replicas) *result {
	// code here
	wg := sync.WaitGroup{}
	ch := make(chan *result, 1)
	ctx2, finish := context.WithCancel(context.TODO())
	defer finish()

	wg.Add(len(replicas))
	for _, srch := range replicas {
		go worker(ctx2, &wg, srch, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	select {
	case <-ctx.Done():
		// fmt.Println("Context finished first")
		return &result{
			err: ctx.Err(),
		}
	case value := <-ch:
		// fmt.Println("Result gained first")
		return value
	}
}

/*
for each replicaKinds we should run a getFirstResult concurrently,
should return a result for each of them
*/
func getResults(ctx context.Context, replicaKinds []replicas) []*result {
	// code here
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	dataSize := len(replicaKinds)
	results := make([]*result, 0, dataSize)

	wg.Add(len(replicaKinds))
	for _, rplc := range replicaKinds {
		go func() {
			mu.Lock()
			results = append(results, getFirstResult(ctx, rplc))
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return results
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 50*time.Millisecond)
	replicaKinds := []replicas{
		replicas{
			fakeSearch("web1"),
			fakeSearch("web2"),
			fakeSearch("web3"),
		},
		replicas{
			fakeSearch("image1"),
			fakeSearch("image2"),
		},
		replicas{
			fakeSearch("video1"),
			fakeSearch("video2"),
			fakeSearch("video3"),
			fakeSearch("video4"),
		},
	}
	for _, res := range getResults(ctx, replicaKinds) {
		fmt.Println(res.msg, res.err)
	}
}
