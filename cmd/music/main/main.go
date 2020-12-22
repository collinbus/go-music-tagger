package main

import (
	"fmt"
	"go-music-tagger/pkg/music"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide an input argument")
	}
	tag := music.ReadTagFrom(os.Args[1])
	printTags(tag)
	tag.TrackNumber = 5
	music.WriteFileFrom(os.Args[1], tag)
	newTag := music.ReadTagFrom(os.Args[1])
	printTags(newTag)
}

func printTags(tag *music.Tag) {
	fmt.Println(tag.Title)
	fmt.Println(tag.Artists[0])
	fmt.Println(tag.Album)
	fmt.Println(tag.Year)
	fmt.Println(tag.Genre)
	fmt.Println(tag.TrackNumber)
	fmt.Println(tag.Isrc)
}
