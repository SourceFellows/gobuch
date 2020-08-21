package main

import (
	"encoding/json"

	"golang.source-fellows.com/samples/packagename/util"
)

func main() {

	util.ParseValue("Hello World")
	json.Marshal("Hello World")

}
