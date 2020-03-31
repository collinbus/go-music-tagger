package flac

import (
	"io/ioutil"
	"testing"
)

func readSeekTableDataFromFile() (*SeekTable, []byte) {
	seekTable := &SeekTable{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := BlockInfo{length: 558, startIndex: 46, isLastBlock: false}
	seekTable.BlockInfo = blockInfo
	return seekTable, data[46:604]
}

func TestNumberOfSeekPointsInSeekTable(t *testing.T) {
	expectedNumberOfSeekPoints := 31
	seektTable, data := readSeekTableDataFromFile()

	seektTable.Read(data)

	if seektTable.NumberOfSeekPoints != expectedNumberOfSeekPoints {
		t.Errorf("Expected number of seek points %d, but was %d", expectedNumberOfSeekPoints, seektTable.NumberOfSeekPoints)
	}
}

func TestFirstSeekPointInSeekTable(t *testing.T) {
	expectedFirstSampleNumber := uint64(0)
	expectedOffset := uint64(0)
	expectedNumberOfTargetFrames := uint16(4096)
	seekTable, data := readSeekTableDataFromFile()

	seekTable.Read(data)

	seekPoint := seekTable.SeekPoints[0]
	if seekPoint.FirstSampleNumber != expectedFirstSampleNumber {
		t.Errorf("Expected first sample number %d, but was %d", expectedFirstSampleNumber, seekPoint.FirstSampleNumber)
	}
	if seekPoint.Offset != expectedOffset {
		t.Errorf("Expected offset %d, but was %d", expectedOffset, seekPoint.Offset)
	}
	if seekPoint.NumberOfTargetFrames != expectedNumberOfTargetFrames {
		t.Errorf("Expected number of target frames %d, but was %d", expectedNumberOfTargetFrames, seekPoint.NumberOfTargetFrames)
	}
}

func TestThirdSeekPointInSeekTable(t *testing.T) {
	expectedFirstSampleNumber := uint64(880640)
	expectedOffset := uint64(1889510)
	expectedNumberOfTargetFrames := uint16(4096)
	seekTable, data := readSeekTableDataFromFile()

	seekTable.Read(data)

	seekPoint := seekTable.SeekPoints[2]
	if seekPoint.FirstSampleNumber != expectedFirstSampleNumber {
		t.Errorf("Expected first sample number %d, but was %d", expectedFirstSampleNumber, seekPoint.FirstSampleNumber)
	}
	if seekPoint.Offset != expectedOffset {
		t.Errorf("Expected offset %d, but was %d", expectedOffset, seekPoint.Offset)
	}
	if seekPoint.NumberOfTargetFrames != expectedNumberOfTargetFrames {
		t.Errorf("Expected number of target frames %d, but was %d", expectedNumberOfTargetFrames, seekPoint.NumberOfTargetFrames)
	}
}
