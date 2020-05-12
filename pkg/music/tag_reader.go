package music

import (
	"go-music-tagger/internal/flac"
	"log"
	"os"
	"path/filepath"
)

const flacFileExtension = "flac"

func ReadTagFrom(path string) *Tag {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatal(err)
	}

	if filepath.Ext(path) == flacFileExtension {
		return readFlacFile(path)
	}
	return nil
}

func readFlacFile(file string) *Tag {
	fileService := flac.NewFileService()
	flacFile, err := fileService.ReadFile(file)

	if err != nil {
		log.Fatal(err)
	}

	return readTagFrom(flacFile)
}
