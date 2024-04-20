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

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":80", r)
}
