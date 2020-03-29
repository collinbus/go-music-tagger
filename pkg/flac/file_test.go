package flac

import "testing"

const filePath = "../../assets/clocks.flac"
const lyricsFilePath = "../../assets/lyrics.txt"
const emptyFilePath = "../../assets/empty.txt"

var streamInfoLength = 0

func TestReadFileShouldReturnCorrectFileSizeAndBytes(t *testing.T) {
	expectedFileSize := 35804910
	fileReader := NewFileReader(MockStreamInfo{})

	file, _ := fileReader.ReadFile(filePath)

	if file.Size != expectedFileSize {
		t.Errorf("File size should be %d but was %d", expectedFileSize, file.Size)
	}
}

func TestReadFileShouldReturnErrorWhenPathIsInvalid(t *testing.T) {
	wrongPath := "a wrong path"
	expectedError := "open a wrong path: The system cannot find the file specified."
	fileReader := NewFileReader(MockStreamInfo{})

	_, err := fileReader.ReadFile(wrongPath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldFailIfFileIsNotFlacFile(t *testing.T) {
	expectedError := "file at " + lyricsFilePath + " is not a flac file"
	fileReader := NewFileReader(MockStreamInfo{})

	_, err := fileReader.ReadFile(lyricsFilePath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldFailIfFileIsTooSmall(t *testing.T) {
	expectedError := "file at " + emptyFilePath + " is not a flac file"
	fileReader := NewFileReader(MockStreamInfo{})

	_, err := fileReader.ReadFile(emptyFilePath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldReadStreamInfoCorrectly(t *testing.T) {
	fileReader := NewFileReader(MockStreamInfo{})

	file, _ := fileReader.ReadFile(filePath)

	if file.StreamInfo == nil {
		t.Error("Stream info should not be nil")
	}
}

func TestPassCorrectNumberOfBytesToStreamInfoReader(t *testing.T) {
	streamInfoLength = 0
	expectedNumberOfBytes := 34
	fileReader := NewFileReader(MockStreamInfo{})

	_, _ = fileReader.ReadFile(filePath)

	if streamInfoLength != expectedNumberOfBytes {
		t.Errorf("Expected number of bytes passed to stream info reader %d, but was %d", expectedNumberOfBytes, streamInfoLength)
	}
}

type MockStreamInfo struct{}

func (mock MockStreamInfo) NewStreamInfo(streamInfoData []byte) StreamInfo {
	streamInfoLength = len(streamInfoData)
	return StreamInfo{}
}