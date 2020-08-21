package main

import (
	"net/http"
	"time"
)

func main() {
	client := http.Client{Transport: &http.Transport{
		ResponseHeaderTimeout: 1 * time.Millisecond,
	}}
	res, err := client.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
}
