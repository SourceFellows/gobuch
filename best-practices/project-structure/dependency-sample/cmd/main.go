package main

import (
	"net/http"

	"golang.source-fellows.com/samples/applicationx/http/rest"
	"golang.source-fellows.com/samples/applicationx/postgres"
)

func main() {

	us := &postgres.UserService{}

	http.HandleFunc("/", rest.Handler(us))
	http.ListenAndServe(":8080", nil)

}
