package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	certFile := "./certs/test.example.crt"
	keyFile := "./certs/test.example.key"

	http.ListenAndServeTLS(":8443", certFile, keyFile, nil)

}
