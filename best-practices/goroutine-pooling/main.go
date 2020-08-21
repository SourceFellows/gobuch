package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	concurrent = 5
	pool       = make(chan interface{}, concurrent)
)

func work(num int, w *sync.WaitGroup) {
	pool <- "starting"
	defer func() {
		w.Done()
		<-pool
	}()
	fmt.Printf("number %d\n", num)
	time.Sleep(1 * time.Second)
}

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()

}
