package flac

import (
	"errors"
	"io/ioutil"
)

var flacFileIdentifier = []byte{0x66, 0x4C, 0x61, 0x43}

type FileReader struct{}

type File struct {
	Size int
}

func (FileReader) ReadFile(path string) (*File, error) {
	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if !isFlacFile(fileBytes) {
		return nil, errors.New("file at " + path + " is not a flac file")
	}

	return &File{
		Size: len(fileBytes),
	}, nil
}

func isFlacFile(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	for i := range flacFileIdentifier {
		if data[i] != flacFileIdentifier[i] {
			return false
		}
	}
	return true
}

func NewFileReader() *FileReader {
	return &FileReader{}
}
