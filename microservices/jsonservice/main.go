package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID        int    `json:"id,omitempty"`
	Firstname string `json:"first,omitempty"`
	Lastname  string `json:"last,omitempty"`
}

var customer *Customer

func myGetFunc(rw http.ResponseWriter, rq *http.Request) {
	bites, _ := json.Marshal(customer)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(bites)
}

func myPostFunc(rw http.ResponseWriter, rq *http.Request) {
	body, _ := ioutil.ReadAll(rq.Body)
	json.Unmarshal(body, customer)
	rw.WriteHeader(204)
}

func main() {
	customer = &Customer{ID: 1, Firstname: "Peter"}
	r := mux.NewRouter()
	r.HandleFunc("/customer", myGetFunc).Methods("GET")
	r.HandleFunc("/customer", myPostFunc).Methods("POST")
	http.ListenAndServe(":8080", r)
}
