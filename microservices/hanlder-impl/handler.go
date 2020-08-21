package httphandler

import "net/http"

type MyHandler struct{}

func (mh *MyHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	writer.Write([]byte("Hello World!"))
}
