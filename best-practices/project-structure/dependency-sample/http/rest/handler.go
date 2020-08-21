package rest

import (
	"fmt"
	"net/http"
	"time"

	"golang.source-fellows.com/samples/applicationx"
)

func Handler(us applicationx.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		us.CreateUser(&applicationx.User{})
		fmt.Fprintf(w, "Hello World! %s", time.Now())
	}
}
