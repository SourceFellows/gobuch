package main

import (
	"net/http"

	libredis "github.com/go-redis/redis/v7"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/redis"
)

func handleRequest(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World!"))
}

func rateLimitMiddleWare(handler http.HandlerFunc) http.Handler {
	rate, err := limiter.NewRateFromFormatted("4-S")
	if err != nil {
		panic(err)
	}
	// Redis Client anlegen
	option, err := libredis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}
	client := libredis.NewClient(option)
	// Store mit Redis konfigurieren
	store, err := redis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_http_example",
		MaxRetry: 3,
	})
	if err != nil {
		panic(err)
	}
	middleware := stdlib.NewMiddleware(limiter.New(store, rate, limiter.WithTrustForwardHeader(true)))
	return middleware.Handler(handler)
}

func main() {
	http.Handle("/", rateLimitMiddleWare(http.HandlerFunc(handleRequest)))
	http.ListenAndServe(":8080", nil)
}
