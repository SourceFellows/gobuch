package main

import (
	"encoding/json"
	"net/http"

	"github.com/moogar0880/problems"
)

func handleHttp(res http.ResponseWriter, req *http.Request) {
	NotFound := problems.NewStatusProblem(http.StatusNotFound)
	bites, _ := json.Marshal(NotFound)
	res.WriteHeader(http.StatusNotFound)
	res.Header().Add("Content-Type", "application/json")
	res.Write(bites)

}

func main() {
	http.HandleFunc("/a/", handleHttp)
	http.ListenAndServe(":8080", nil)
}
