package flac

import (
	"io/ioutil"
	"testing"
)

func readStreamInfoFromFile() []byte {
	file, _ := ioutil.ReadFile(filePath)
	return file[8:42]
}

func TestReadSampleDataFromFlacStreamInfo(t *testing.T) {
	expectedBlockSize := uint16(4096)
	data := readStreamInfoFromFile()
	info := NewStreamInfo()

	info.readStreamInfo(data)

	if info.MinimumSampleBlockSize != expectedBlockSize {
		t.Errorf("Expected minimum sample block size %d, but was %d", expectedBlockSize, info.MinimumSampleBlockSize)
	}
	if info.MaximumSampleBlockSize != expectedBlockSize {
		t.Errorf("Expected maximum sample block size %d, but was %d", expectedBlockSize, info.MaximumSampleBlockSize)
	}
}
