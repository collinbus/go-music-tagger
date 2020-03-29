package flac

import (
	"io/ioutil"
)

type File struct {
	Size int
}

func ReadFile(path string) (*File, error) {
	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	return &File{
		Size: len(bytes),
	}, nil
}
