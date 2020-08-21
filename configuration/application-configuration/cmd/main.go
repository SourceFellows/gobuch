/*Package main show how to use envconfig.

Start the app in the root folder with:

	SERVER_PORT=8080 SERVER_HOST=localhost go run cmd/main.go

*/
package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.source-fellows.com/samples/applicationconfiguration"
)

func main() {

	cfg := applicationconfiguration.Config{}
	err := envconfig.Process("", &cfg)

	if err != nil {
		panic(err)
	}

	fmt.Println(cfg.Server.Port)

}
