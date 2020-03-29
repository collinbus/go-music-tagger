package flac

import "testing"

const filePath = "../../assets/clocks.flac"

func TestReadFileShouldReturnCorrectFileSizeAndBytes(t *testing.T) {
	expectedFileSize := 35804910

	file := ReadFile(filePath)

	if file.Size != expectedFileSize {
		t.Errorf("File size should be %d but was %d", expectedFileSize, file.Size)
	}
}
