package main

import (
	"fmt"
	"go-music-tagger/internal/flac"
)

func main() {
	fmt.Println("Go Music Tagger")
	service := flac.NewFileService()
	file, _ := service.ReadFile("C:\\Users\\Collin\\go\\src\\go-music-tagger\\assets\\clocks.flac")
	_ = flac.WriteFile(*file, "C:\\Users\\Collin\\go\\src\\go-music-tagger\\assets\\new-clocks.flac")
}
