package main

import "net/http"

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello World (from ServeHTTP)"))
}

func handleHttp(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Hello World (from handleHttp)"))
}

func main() {
	//Anmeldung der HandlerFunc Implementierung
	http.HandleFunc("/func", handleHttp)
	//Anmeldung der Handler Interface Implementierung
	http.Handle("/handler", &MyHandler{})
	http.ListenAndServe(":8080", &MyHandler{})
}
