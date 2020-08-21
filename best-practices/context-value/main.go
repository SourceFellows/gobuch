package main

import (
	"context"
	"fmt"
)

type contextKey string

func main() {
	var key contextKey = "test"
	ctx := context.Background()
	ctx = context.WithValue(ctx, key, "world")
	fmt.Println(ctx.Value(key))
}
