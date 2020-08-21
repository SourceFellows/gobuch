package main

import (
	"log"

	"github.com/hashicorp/go-retryablehttp"
)

func main() {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	standardClient := retryClient.StandardClient()

	res, err := standardClient.Get("http://localhost:8080/")
	if err != nil {
		log.Fatal("could not request")
	}
	defer res.Body.Close()
	log.Println(res.Status)
}
