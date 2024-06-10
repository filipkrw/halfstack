package main

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"os"

	"halfstack/components"

	"github.com/a-h/templ"
	"github.com/gorilla/mux"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/parser"
)

func static() {
	f, err := os.Create("hello.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = components.Homepage().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}

func main() {
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
