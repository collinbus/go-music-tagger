package flac

import (
	"testing"
)

const filePath = "../../assets/clocks.flac"
const lyricsFilePath = "../../assets/lyrics.txt"
const emptyFilePath = "../../assets/empty.txt"

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

func TestPassCorrectSizeStartAndLastBlockInfoToStreamInfoReader(t *testing.T) {
	expectedNumberOfBytes := uint32(34)
	expectedStartIndex := 8
	fileReader := NewFileReader(MockStreamInfo{})

	file, _ := fileReader.ReadFile(filePath)

	blockLength := file.StreamInfo.BlockInfo.Length
	start := file.StreamInfo.BlockInfo.StartIndex
	isLastBlock := file.StreamInfo.BlockInfo.isLastBlock
	if blockLength != expectedNumberOfBytes {
		t.Errorf("Expected size of stream info %d, but was %d", expectedNumberOfBytes, blockLength)
	}
	if start != expectedStartIndex {
		t.Errorf("Expected start index of stream info %d, but was %d", expectedNumberOfBytes, start)
	}
	if isLastBlock {
		t.Error("StreamInfo should not be the latest block of the metadata")
	}
}

func TestPassCorrectSizeStartAndLastBlockInfoToSeekTable(t *testing.T) {
	expectedNumberOfBytes := uint32(558)
	expectedStartIndex := 46
	fileReader := NewFileReader(MockStreamInfo{})

	file, _ := fileReader.ReadFile(filePath)

	blockLength := file.SeekTable.BlockInfo.Length
	start := file.SeekTable.BlockInfo.StartIndex
	isLastBlock := file.SeekTable.BlockInfo.isLastBlock
	if blockLength != expectedNumberOfBytes {
		t.Errorf("Expected size of seek table %d, but was %d", expectedNumberOfBytes, blockLength)
	}
	if start != expectedStartIndex {
		t.Errorf("Expected start index of seek table %d, but was %d", expectedNumberOfBytes, start)
	}
	if isLastBlock {
		t.Error("SeekTable should not be the latest block of the metadata")
	}
}

type MockStreamInfo struct{}

func (mock MockStreamInfo) Read(_ []byte, _ interface{}) {

}
