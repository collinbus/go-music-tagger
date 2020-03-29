package flac

import (
	"io/ioutil"
	"log"
)

type File struct {
	Size int
}

func ReadFile(path string) File {
	bytes, _ := ioutil.ReadFile(path)

	return File{
		Size: len(bytes),
	}
}
