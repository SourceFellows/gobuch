package main

import "net/http"

func handleForm(rw http.ResponseWriter, rq *http.Request) {
	err := rq.ParseForm()
	if err != nil {
		rw.WriteHeader(http.StatusNotAcceptable)
		return
	}
	name := rq.Form.Get("name")
	rw.Write([]byte("Hello " + name))
}

func main() {
	http.HandleFunc("/", handleForm)
	http.ListenAndServe(":8080", nil)
}
