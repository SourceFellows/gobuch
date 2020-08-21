package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	certPool := x509.NewCertPool()
	crt, _ := ioutil.ReadFile("./certs/myRoot.crt")
	certPool.AppendCertsFromPEM(crt)
	keyPair, _ := tls.LoadX509KeyPair("./certs/user.crt", "./certs/user.key")
	clientCerts := []tls.Certificate{keyPair}
	tlsConfig := &tls.Config{
		RootCAs:      certPool,
		Certificates: clientCerts,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := http.Client{Transport: transport}
	res, err := client.Get("https://test.example.private:8443")
	if err != nil {
		panic(err)
	}
	bites, err := ioutil.ReadAll(res.Body)
	fmt.Printf("response is: %v", string(bites))
}
