package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
)

type Result struct {
	Name string
}

func myHandler(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if len(path) < 6 || path[:6] != "/json/" {
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}
	res := &Result{Name: path[6:]}
	bites, err := json.Marshal(*res)
	if err != nil {
		writer.WriteHeader(500)
		return
	}
	writer.Header().Add("Content-Type", "application/json")
	writer.Write(bites)
}

func main() {
	http.HandleFunc("/", myHandler)
	http.ListenAndServe(":8080", nil)
}
