package main

import (
	"context"
	"halfstack/components"
	"log"
	"os"
)

func generate() {
	f, err := os.Create("dist/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = components.Homepage().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
