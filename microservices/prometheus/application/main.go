package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(rs http.ResponseWriter, rq *http.Request) {
		rs.WriteHeader(http.StatusOK)
	})
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
