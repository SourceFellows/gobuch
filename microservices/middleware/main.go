package main

import (
	"log"
	"net/http"
)

func businessLogic(text string) http.Handler {
	myFunc := func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(text))
	}
	return http.HandlerFunc(myFunc)
}

func middleWare(wrapped http.Handler) http.Handler {
	return http.HandlerFunc(func(re http.ResponseWriter, rq *http.Request) {
		log.Println("before request")
		wrapped.ServeHTTP(re, rq)
		log.Println("after request")
	})
}

func main() {
	http.Handle("/", middleWare(businessLogic("Hello World")))
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Printf("error with server %v", err)
	}
}
