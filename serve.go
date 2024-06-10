package main

import (
	"bytes"
	"halfstack/components"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func serve() {
	r := mux.NewRouter()

	r.Handle("/", templ.Handler(components.Homepage()))

	md := goldmark.New(
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
	)

	source, _ := os.ReadFile("posts/post.md")

	var buf bytes.Buffer

	if err := md.Convert(source, &buf); err != nil {
		panic(err)
	}

	r.Handle("/post", templ.Handler(components.Post(buf.String())))

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.ListenAndServe(":80", r)
}
