package flac

import (
	"io/ioutil"
	"testing"
)

func readVorbisCommentDataFromFile() (*VorbisComment, []byte) {
	vorbisComment := &VorbisComment{}
	data, _ := ioutil.ReadFile(filePath)
	blockInfo := &BlockInfo{length: 1813, startIndex: 608, isLastBlock: false}
	vorbisComment.BlockInfo = blockInfo
	return vorbisComment, data[608:1813]
}

func TestVendorInVorbisComment(t *testing.T) {
	expectedVendor := "reference libFLAC 1.2.1 20070917"
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.Vendor != expectedVendor {
		t.Errorf("Expected vendor %s, but is %s", expectedVendor, vorbisComment.Vendor)
	}
}

func TestNumberOfCommentsInVorbisComment(t *testing.T) {
	expectedNumberOfConmments := 12
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.NumberOfComments != expectedNumberOfConmments {
		t.Errorf("Expected number of comments %d, but is %d", expectedNumberOfConmments, vorbisComment.NumberOfComments)
	}
}

func TestNumberFirstCommentInVorbisComment(t *testing.T) {
	expectedComment := "album=A Rush Of Blood To The Head"
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.Comments[0] != expectedComment {
		t.Errorf("Expected comment %s, but is%s", expectedComment, vorbisComment.Comments[0])
	}
}
