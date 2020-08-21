package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	tc := make(chan time.Time)

	var c int64 = 0

	go func() {
		for {
			go func() {
				for {
					c++
					i := rand.Int31n(10)
					time.Sleep(time.Duration(i) * time.Millisecond)
					tc <- time.Now()
				}
			}()
		}
	}()

	for t := range tc {
		fmt.Println(t)
		fmt.Println(c)
	}

}
