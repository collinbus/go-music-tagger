package flac

import (
	"io/ioutil"
	"testing"
)

func readPictureDataFromFile() (*Picture, []byte) {
	picture := &Picture{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := &BlockInfo{length: 81660, startIndex: 0, isLastBlock: false}
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

func TestLengthOfPicturesInBytes(t *testing.T) {
	expectedLengthOfPictureData := 9883
	picture, data := readPictureDataFromFile()

	picture.Read(data)

	pictureData := picture.PictureData
	if len(pictureData) != expectedLengthOfPictureData {
		t.Errorf("Expected picture data length %d, but was %d", expectedLengthOfPictureData, len(pictureData))
	}
}

func TestWritePicture(t *testing.T) {
	expectedBytes := []byte{0x06, 0x0, 0x0, 0x3a, 0x0, 0x0, 0x0, 0x03, 0x0, 0x0, 0x0, 0x0b, 0x61, 0x5f, 0x6d, 0x69, 0x6d, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
		0x0, 0x0, 0x0, 0x0d, 0x61, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x0, 0x0, 0x0, 0x05, 0x0, 0x0, 0x0, 0x0a,
		0x0, 0x0, 0x0, 0x02, 0x0, 0x0, 0x0, 0x03, 0x0, 0x0, 0x0, 0x02, 0x05, 0xe6}

	picture := NewPicture(NewBlockInfo(8, 58, false))
	picture.PictureType = 3
	picture.MimeType = "a_mime_type"
	picture.Description = "a_description"
	picture.Width = 5
	picture.Height = 10
	picture.ColorDepth = 2
	picture.IndexedColorPictures = 3
	picture.PictureData = []byte{0x05, 0xe6}

	data := picture.WritePicture()

	if data == nil {
		t.Errorf("The data should not be nil")
	}

	if len(data) != 62 {
		t.Errorf("The binary data should contain 62 bytes but it contains %d", len(data))
		return
	}

	for i, b := range data {
		if b != expectedBytes[i] {
			t.Errorf("byte at index %d should be %d but was %d", i, expectedBytes[i], b)
		}
	}
}
