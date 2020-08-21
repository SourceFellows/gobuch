package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println(res.Status)
	bites, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(bites))
}
