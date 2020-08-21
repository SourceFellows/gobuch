package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	bg := context.Background()
	fmt.Println(bg.Deadline()) //0001-01-01 00:00:00 +0000 UTC false
	fmt.Println(bg.Err())      //<nil>
	//fmt.Println(<-td.Done())   //blockiert für immer!

	ctx, cancel := context.WithTimeout(bg, 5*time.Second)
	defer cancel()

	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel()
	}()

	fmt.Println(ctx.Deadline())
	<-ctx.Done()
	fmt.Println("Ende")

}
