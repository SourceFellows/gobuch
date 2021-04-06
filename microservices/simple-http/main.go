package main

import "net/http"

func handleHttp(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World"))
}

func main() {
	http.HandleFunc("/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
