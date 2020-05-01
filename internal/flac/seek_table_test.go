package flac

import (
	"io/ioutil"
	"testing"
)

func readSeekTableDataFromFile() (*SeekTable, []byte) {
	seekTable := &SeekTable{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := &BlockInfo{length: 558, startIndex: 46, isLastBlock: false}
	seekTable.BlockInfo = blockInfo
	return seekTable, data[46:604]
}

func TestNumberOfSeekPointsInSeekTable(t *testing.T) {
	expectedNumberOfSeekPoints := 31
	seekTable, data := readSeekTableDataFromFile()

	seekTable.Read(data)

	if seekTable.NumberOfSeekPoints != expectedNumberOfSeekPoints {
		t.Errorf("Expected number of seek points %d, but was %d", expectedNumberOfSeekPoints, seekTable.NumberOfSeekPoints)
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

func TestWriteSeekTable(t *testing.T) {
	expectedBytes := []byte{0x03, 0x00, 0x00, 0x24, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10, 0x00, 0x32,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x56, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x13, 0x04, 0x00, 0x2c,
	}
	seekPoints := []SeekPoint{{
		FirstSampleNumber:    8,
		Offset:               16,
		NumberOfTargetFrames: 50,
	}, {
		FirstSampleNumber:    86,
		Offset:               4868,
		NumberOfTargetFrames: 44,
	}}
	seekTable := &SeekTable{
		BlockInfo:          &BlockInfo{startIndex: 4, length: 36, isLastBlock: false},
		NumberOfSeekPoints: 2,
		SeekPoints:         seekPoints,
	}

	points := seekTable.WriteSeekPoints()

	if points == nil {
		t.Errorf("Seektable bytes should not be nil")
		return
	}

	if points[3] != 0x24 {
		t.Errorf("Block length should be %d but was %d", 36, points[3])
		return
	}

	for i, b := range points {
		if b != expectedBytes[i] {
			t.Errorf("byte at index %d should be %d but was %d", i, expectedBytes[i], b)
		}
	}
}
