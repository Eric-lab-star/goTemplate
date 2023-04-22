package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe("localhost:5000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "home")
}
