package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	postCallsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_post_total",
		Help: "The total number of POST requests",
	})
	getCallsCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_get_total",
		Help: "The total number of GET requests",
	})
)

func main() {
	http.HandleFunc("/do", func(rs http.ResponseWriter, rq *http.Request) {

		if rq.Method == http.MethodPost {
			postCallsCounter.Add(1)
		} else if rq.Method == http.MethodGet {
			getCallsCounter.Add(1)
		}

		fmt.Fprintln(rs, "Hello World")
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
