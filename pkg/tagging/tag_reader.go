package tagging

import (
	"go-music-tagger/internal/flac"
	"log"
	"os"
	"path/filepath"
)

const flacFileExtension = ".flac"

func ReadTagFrom(musicFile *os.File) *Tag {
	if filepath.Ext(musicFile.Name()) == flacFileExtension {
		return readFlacFile(musicFile.Name())
	}
	return nil
}

func readFlacFile(file string) *Tag {
	flacFile, err := flac.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return readTagFrom(flacFile)
}
