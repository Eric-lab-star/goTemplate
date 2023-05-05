package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.Handle("/folder/", http.StripPrefix("/folder", http.FileServer(http.Dir("static"))))
	http.ListenAndServe("localhost:4000", nil)
}

type homeHandler func(w http.ResponseWriter, r *http.Request)

func (f homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home ")
}
