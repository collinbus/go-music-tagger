package flac

import "testing"

func TestWriteMetaDataBlockHeader(t *testing.T) {
	expectedHeader := []byte{0x00, 0x00, 0x00, 0x22}
	header := WriteBlockHeader(false, StreamInfoBlock, 34)

	if header == nil {
		t.Errorf("The header should never be nil")
	}

	for i, b := range header {
		if b != expectedHeader[i] {
			t.Errorf("byte at index %d should be %d but was %d", i, expectedHeader[i], b)
		}
	}
}

func TestWriteMetaDataBlockHeaderWithLastBlock(t *testing.T) {
	expectedHeader := []byte{0x83, 0x00, 0x00, 0x22}
	header := WriteBlockHeader(true, SeekTableBlock, 34)

	if header == nil {
		t.Errorf("The header should never be nil")
	}

	for i, b := range header {
		if b != expectedHeader[i] {
			t.Errorf("byte at index %d should be %d but was %d", i, expectedHeader[i], b)
		}
	}
}
