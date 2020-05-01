package flac

import (
	"bytes"
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

func TestNumberFourthCommentInVorbisComment(t *testing.T) {
	file, _ := ioutil.ReadFile(lyricsFilePath)
	expectedComment := bytes.NewBuffer(file).String()
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.Comments[3] != expectedComment {
		t.Errorf("Expected comment: \n%s\nbut is:\n%s", expectedComment, vorbisComment.Comments[3])
	}
}

func TestFifthCommentInVorbisComment(t *testing.T) {
	expectedComment := "artist=Coldplay"
	vorbisComment, data := readVorbisCommentDataFromFile()

	vorbisComment.Read(data)

	if vorbisComment.Comments[4] != expectedComment {
		t.Errorf("Expected comment %s, but is %s", expectedComment, vorbisComment.Comments[4])
	}
}

func TestWriteVorbisComment(t *testing.T) {
	expectedBytes := []byte{0x04, 0x00, 0x00, 0x3f, 0x16, 0x00, 0x00, 0x00, 0x67, 0x6f, 0x2d, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x2d, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x20, 0x76, 0x30, 0x2e, 0x30, 0x2e, 0x31,
		0x02, 0x00, 0x00, 0x00, 0x0b, 0x00, 0x00, 0x00, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x31, 0x3d, 0x68, 0x69,
		0x0e, 0x00, 0x00, 0x00, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x32, 0x3d, 0x68, 0x65, 0x6c, 0x6c, 0x6f}
	vc := NewVorbisComment(NewBlockInfo(4, 22, false))
	vc.Vendor = "go-music-tagger v0.0.1"
	vc.NumberOfComments = 2
	vc.Comments = []string{"comment1=hi", "comment2=hello"}

	data := vc.WriteVorbisComments()

	if data == nil {
		t.Errorf("The data should not be nil")
	}

	for i, b := range data {
		if b != expectedBytes[i] {
			t.Errorf("byte at index %d should be %d but was %d", i, expectedBytes[i], b)
		}
	}
}
