package main

import (
	"go-music-tagger/internal/flac"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("Please provide an input and output argument")
	}
	file, _ := flac.ReadFile(os.Args[1])
	_ = flac.WriteFile(file, os.Args[2])
}
