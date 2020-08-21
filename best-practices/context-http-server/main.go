package main

import (
	"fmt"
	"net/http"
	"time"
)

func handleCall(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	select {
	case <-ctx.Done():
		res.Write([]byte("end"))
		fmt.Println("clsoing connection to client")
		return
	case <-time.After(20 * time.Second):
		fmt.Fprintln(res, "Hello World")
	}
}

func main() {
	http.HandleFunc("/", handleCall)
	http.ListenAndServe(":8080", nil)
}
