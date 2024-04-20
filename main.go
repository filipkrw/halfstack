package main

import (
	"net/http"

	"halfstack/components"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/", templ.Handler(components.Hello("Filip")))

	http.ListenAndServe(":80", r)
}
