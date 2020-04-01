package flac

import (
	"io/ioutil"
	"testing"
)

func readPictureDataFromFile() (*Picture, []byte) {
	picture := &Picture{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := &BlockInfo{length: 1813, startIndex: 608, isLastBlock: false}
	picture.BlockInfo = blockInfo
	return picture, data[1817:11742]
}

func TestPictureTypeOfPicture(t *testing.T) {
	expectedPictureType := 3
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	if int(picture.PictureType) != expectedPictureType {
		t.Errorf("Expected picture type %d, but was %d", expectedPictureType, picture.PictureType)
	}
}
