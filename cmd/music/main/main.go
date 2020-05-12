package main

import (
	"fmt"
	"go-music-tagger/pkg/music"
)

func main() {
	tag := music.ReadTagFrom("C:\\Users\\Collin\\go\\src\\go-music-tagger\\assets\\new-clocks.flac")
	fmt.Println(tag.Title)
	fmt.Println(tag.Artists[0])
	fmt.Println(tag.Album)
	fmt.Println(tag.Year)
	fmt.Println(tag.Genre)
	fmt.Println(tag.TrackNumber)
	fmt.Println(tag.Isrc)
}
