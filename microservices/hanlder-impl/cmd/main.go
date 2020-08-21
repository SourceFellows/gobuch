package main

import (
	"net/http"

	"golang.source-fellows.com/samples/httphandler"
)

func main() {
	http.ListenAndServe(":8080", &httphandler.MyHandler{})
}
