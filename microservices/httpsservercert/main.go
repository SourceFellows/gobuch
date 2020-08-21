package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	})

	certPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile("certs/server-chain.crt")
	if err != nil {
		panic(err)
	}
	certPool.AppendCertsFromPEM(pem)

	tlsConfig := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  certPool,
	}

	srv := http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	log.Fatal(srv.ListenAndServeTLS("certs/test.example.crt", "certs/test.example.key"))
}
