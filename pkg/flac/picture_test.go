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

func TestMimeTypeOfPicture(t *testing.T) {
	expectedMimeType := "image/jpeg"
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	if picture.MimeType != expectedMimeType {
		t.Errorf("Expected mime type %s, but was %s", expectedMimeType, picture.MimeType)
	}
}

func TestDescriptionOfPicture(t *testing.T) {
	expectedDescription := ""
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	if picture.Description != expectedDescription {
		t.Errorf("Expected mime type %s, but was %s", expectedDescription, picture.Description)
	}
}

func TestWidthAndHeightInPixelsOfPicture(t *testing.T) {
	expectedWidth := uint32(0)
	expectedHeight := uint32(0)
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	if picture.Width != expectedWidth {
		t.Errorf("Expected width %d, but was %d", expectedWidth, picture.Width)
	}
	if picture.Height != expectedHeight {
		t.Errorf("Expected height %d, but was %d", expectedHeight, picture.Height)
	}
}

func TestColorDepthAndIndexedColorPicturesOfPicture(t *testing.T) {
	expectedColorDepth := uint32(0)
	expectedIndexedColorPictures := uint32(0)
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	if picture.ColorDepth != expectedColorDepth {
		t.Errorf("Expected colordepth %d, but was %d", expectedColorDepth, picture.ColorDepth)
	}
	if picture.IndexedColorPictures != expectedIndexedColorPictures {
		t.Errorf("Expected indexed color pictures %d, but was %d", expectedIndexedColorPictures, picture.IndexedColorPictures)
	}
}
