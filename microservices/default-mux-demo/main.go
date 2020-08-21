package main

import "net/http"

type myHandler struct {
	Name string
}

func (h *myHandler) ServeHTTP(r http.ResponseWriter, _ *http.Request) {
	r.Write([]byte(h.Name))
}

func main() {
	handler1 := &myHandler{Name: "handler1"}
	handler2 := &myHandler{Name: "handler2"}
	handler3 := &myHandler{Name: "handler3"}
	handler4 := &myHandler{Name: "handler4"}
	handler5 := &myHandler{Name: "handler5"}
	handler6 := &myHandler{Name: "handler6"}

	http.Handle("/", handler1)
	http.Handle("/book", handler2)
	http.Handle("/books/", handler3)
	http.Handle("/books/Hobbit", handler4)
	http.Handle("/theme", handler5)
	http.Handle("/theme/", handler6)
	http.ListenAndServeTLS(":8080", nil)
}
