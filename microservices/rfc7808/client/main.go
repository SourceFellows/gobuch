package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/moogar0880/problems"
)

func main() {
	res, err := http.Get("http://localhost:8080/a")
	if err != nil {
		log.Fatal("could not request")
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		problem := &problems.DefaultProblem{}
		bites, _ := ioutil.ReadAll(res.Body)
		json.Unmarshal(bites, problem)
		log.Println(problem.ProblemTitle())
	} else {
		log.Println("ok")
	}
}
