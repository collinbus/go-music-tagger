package music

import (
	"go-music-tagger/internal/flac"
	"log"
	"os"
	"path/filepath"
)

func WriteFileFrom(originalFile string, tag *Tag) bool {
	if _, err := os.Stat(originalFile); os.IsNotExist(err) {
		log.Fatal(err)
	}

	if filepath.Ext(originalFile) == flacFileExtension {
		return writeFlacFile(originalFile, tag)
	}
	return false
}

func writeFlacFile(originalFile string, tag *Tag) bool {
	file, err := flac.ReadFile(originalFile)
	if err != nil {
		log.Fatal(err)
	}
	newFile := writeFlacTag(file, tag)
	_ = flac.WriteFile(newFile, originalFile)
	return true
}
