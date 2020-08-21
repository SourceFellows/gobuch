package main

import (
	"net/http"
)

var jsonDoc1 = []byte("{\"Name\":\"")
var jsonDoc2 = []byte("\"}")

func myHandler(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if len(path) < 6 || path[:6] != "/json/" {
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}

	writer.Header().Add("Content-Type", "application/json")
	bites := make([]byte, 20)
	bites = append(bites, jsonDoc1...)
	bites = append(bites, path[6:]...)
	bites = append(bites, jsonDoc2...)
	writer.Write(bites)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
