package flac

import (
	"testing"
)

const filePath = "../../assets/clocks.flac"
const lyricsFilePath = "../../assets/lyrics.txt"
const emptyFilePath = "../../assets/empty.txt"

func TestReadFileShouldReturnCorrectFileSizeAndBytes(t *testing.T) {
	expectedFileSize := 35804910
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	if file.Size != expectedFileSize {
		t.Errorf("File size should be %d but was %d", expectedFileSize, file.Size)
	}
}

func TestReadFileShouldReturnErrorWhenPathIsInvalid(t *testing.T) {
	wrongPath := "a wrong path"
	expectedError := "open a wrong path: The system cannot find the file specified."
	fileReader := NewFileService()

	_, err := fileReader.ReadFile(wrongPath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldFailIfFileIsNotFlacFile(t *testing.T) {
	expectedError := "file at " + lyricsFilePath + " is not a flac file"
	fileReader := NewFileService()

	_, err := fileReader.ReadFile(lyricsFilePath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldFailIfFileIsTooSmall(t *testing.T) {
	expectedError := "file at " + emptyFilePath + " is not a flac file"
	fileReader := NewFileService()

	_, err := fileReader.ReadFile(emptyFilePath)

	if err == nil || err.Error() != expectedError {
		t.Errorf("Read file should return error: %s", expectedError)
	}
}

func TestReadFileShouldReadStreamInfoCorrectly(t *testing.T) {
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	if file.StreamInfo == nil {
		t.Error("Stream info should not be nil")
	}
}

func TestPassCorrectSizeStartAndLastBlockInfoToStreamInfoReader(t *testing.T) {
	expectedNumberOfBytes := uint32(34)
	expectedStartIndex := 8
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	blockLength := file.StreamInfo.BlockInfo.length
	start := file.StreamInfo.BlockInfo.startIndex
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
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	blockLength := file.SeekTable.BlockInfo.length
	start := file.SeekTable.BlockInfo.startIndex
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

func TestPassCorrectSizeStartAndLastBlockInfoToVorbisComment(t *testing.T) {
	expectedNumberOfBytes := uint32(1205)
	expectedStartIndex := 608
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	blockLength := file.VorbisComment.BlockInfo.length
	start := file.VorbisComment.BlockInfo.startIndex
	isLastBlock := file.VorbisComment.BlockInfo.isLastBlock
	if blockLength != expectedNumberOfBytes {
		t.Errorf("Expected size of vorbis comment %d, but was %d", expectedNumberOfBytes, blockLength)
	}
	if start != expectedStartIndex {
		t.Errorf("Expected start index of vorbis comment %d, but was %d", expectedNumberOfBytes, start)
	}
	if isLastBlock {
		t.Error("Vorbis Comment should not be the latest block of the metadata")
	}
}

func TestPassCorrectSizeStartAndLastBlockInfoToFirstPicture(t *testing.T) {
	expectedNumberOfBytes := uint32(9925)
	expectedStartIndex := 1817
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	firstPicture := file.Picture[0]
	blockLength := firstPicture.BlockInfo.length
	start := firstPicture.BlockInfo.startIndex
	isLastBlock := firstPicture.BlockInfo.isLastBlock
	if blockLength != expectedNumberOfBytes {
		t.Errorf("Expected size of picture %d, but was %d", expectedNumberOfBytes, blockLength)
	}
	if start != expectedStartIndex {
		t.Errorf("Expected start index of picture %d, but was %d", expectedNumberOfBytes, start)
	}
	if isLastBlock {
		t.Error("picture should not be the latest block of the metadata")
	}
}

func TestAddCorrectNumberOfBytesAsAudioData(t *testing.T) {
	expectedAudioSize := 35757252
	fileReader := NewFileService()

	file, _ := fileReader.ReadFile(filePath)

	audioDataLength := len(file.AudioData)
	if audioDataLength != expectedAudioSize {
		t.Errorf("Expected audio data length %d, but was %d", expectedAudioSize, audioDataLength)
	}
}
