package flac

import (
	"io/ioutil"
	"testing"
)

func readVorbisCommentDataFromFile() (*VorbisComment, []byte) {
	vorbisComment := &VorbisComment{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := &BlockInfo{length: 608, startIndex: 1813, isLastBlock: false}
	vorbisComment.BlockInfo = blockInfo
	return vorbisComment, data[608:1813]
}

func TestVendorInVorbisComment(t *testing.T) {
	expectedVendor := "reference libFLAC 1.2.1 2007"
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.Vendor != expectedVendor {
		t.Errorf("Expected vendor %s, but is %s", expectedVendor, vorbisComment.Vendor)
	}
}
