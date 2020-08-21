package main

import (
	"context"
	"fmt"
)

type contextKey int

const usernameKey contextKey = iota

func businessCall(username string) {
	//.. was wichtiges
	fmt.Println(username)
}

func main() {
	ctx := context.Background()
	vCtx := context.WithValue(ctx, usernameKey, "bob")
	//.. einiges anderes
	businessCall(vCtx.Value(usernameKey).(string))

	//Variante2
	username := "bob"
	//.. einiges anderes
	businessCall(username)

}
