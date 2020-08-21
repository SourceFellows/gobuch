package main

import (
	"fmt"
	"net/http"
)

var count int = 0

func handleHttp(res http.ResponseWriter, req *http.Request) {

	responseStatus := http.StatusInternalServerError
	if count%4 == 0 {
		responseStatus = http.StatusOK
	}
	fmt.Printf("will return %v\n", responseStatus)
	res.WriteHeader(responseStatus)

	count++

}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
