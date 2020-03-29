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

func TestReadFrameSizeFromFlacStreamInfo(t *testing.T) {
	expectedMinimumFrameSize := uint32(14)
	expectedMaximumFrameSize := uint32(13616)
	data := readStreamInfoFromFile()
	info := NewStreamInfo()

	info.readStreamInfo(data)

	if info.MinimumFrameSize != expectedMinimumFrameSize {
		t.Errorf("Expected minimum frame size %d, but was %d", expectedMinimumFrameSize, info.MinimumFrameSize)
	}
	if info.MaximumFrameSize != expectedMaximumFrameSize {
		t.Errorf("Expected maximum frame size %d, but was %d", expectedMaximumFrameSize, info.MaximumFrameSize)
	}
}

func TestReadSampleRateFromFlacStreamInfo(t *testing.T) {
	expectedSampleRate := uint32(44100)
	data := readStreamInfoFromFile()
	info := NewStreamInfo()

	info.readStreamInfo(data)

	if info.SampleRate != expectedSampleRate {
		t.Errorf("Expected minimum frame size %d, but was %d", expectedSampleRate, info.SampleRate)
	}
}

func TestReadNumberOfChannelsAndBitsPerSampleFlacStreamInfo(t *testing.T) {
	expectedNumberOfChannels := uint8(2)
	expectedBitsPerSample := uint8(16)
	data := readStreamInfoFromFile()
	info := NewStreamInfo()

	info.readStreamInfo(data)

	if info.NumberOfChannels != expectedNumberOfChannels {
		t.Errorf("Expected number of channels %d, but was %d", expectedNumberOfChannels, info.NumberOfChannels)
	}
	if info.BitsPerSample != expectedBitsPerSample {
		t.Errorf("Expected bits per sample %d, but was %d", expectedBitsPerSample, info.BitsPerSample)
	}
}
