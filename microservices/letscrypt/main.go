package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	log.Fatal(http.Serve(autocert.NewListener("x.source-fellows.com"), nil))

}
