package main

import (
	"log"
	"net/http"

	"golang.source-fellows.com/samples/jwt/server/jwt"
)

func handle(res http.ResponseWriter, req *http.Request) {
	token, err := jwt.CreateToken()
	if err != nil {
		log.Fatal(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = jwt.ValidateToken(token)
	if err != nil {
		log.Fatal(err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	res.Write(token)
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
