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
