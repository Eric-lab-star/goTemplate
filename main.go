package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()

	http.ListenAndServe("localhost:5000", r)
}
