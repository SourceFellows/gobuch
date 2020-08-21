package main

import "net/http"

func myHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		res.Write([]byte("Hello World"))
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
