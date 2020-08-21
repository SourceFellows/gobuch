package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func myFunc(rw http.ResponseWriter, rq *http.Request) {
	rw.Write([]byte("Hello World"))
}

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/").HandlerFunc(myFunc)
	r.Methods("GET")
	http.ListenAndServe(":8080", r)
}
