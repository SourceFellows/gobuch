package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	breaker "github.com/sony/gobreaker"
)

var cb *breaker.CircuitBreaker

func main() {

	settings := breaker.Settings{Interval: 1 * time.Second, Timeout: 5 * time.Second}
	cb = breaker.NewCircuitBreaker(settings)

	for {
		bites, err := cb.Execute(func() (interface{}, error) {
			res, err := http.Get("http://localhost:8080/a")
			if err != nil {
				return nil, err
			}
			if res.StatusCode == http.StatusOK {

			}
			defer res.Body.Close()
			bites, _ := ioutil.ReadAll(res.Body)
			return bites, nil
		})

		if err != nil {
			log.Printf("error %v\n", err)

		} else {
			fmt.Println(string(bites.([]byte)))
		}

	}

}
