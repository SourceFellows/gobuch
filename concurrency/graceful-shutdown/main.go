package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func handle(w http.ResponseWriter, r *http.Request) {

	select {
	case <-time.After(60 * time.Second):
		// If we receive a message after 2 seconds
		// that means the request has been processed
		// We then write this as the response
		w.Write([]byte("request processed"))
	case <-r.Context().Done():
		// If the request gets cancelled, log it
		// to STDERR
		fmt.Fprint(os.Stderr, "request cancelled\n")
	}

}

func main() {

	srv := http.Server{Addr: ":8081"}

	idleConnsClosed := make(chan struct{})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		log.Print("HTTP server Shutdown successful")
		close(idleConnsClosed)
	}()

	http.HandleFunc("/", handle)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}
