package main

import (
	"encoding/json"
	"net/http"
	_ "net/http/pprof"
	"regexp"
)

var re = regexp.MustCompile("^/json/(.*)$")

type Result struct {
	Name string
}

func myHandler(writer http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if !re.MatchString(path) {
		writer.WriteHeader(http.StatusNotImplemented)
		return
	}
	parts := re.FindStringSubmatch(path)
	res := &Result{Name: parts[1]}
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
