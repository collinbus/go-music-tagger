package music

import (
	"go-music-tagger/internal/flac"
	"testing"
)

var artists = []string{"an_artist", "a_featuring_artist"}

const trackNumber = 1

func TestWriteFlacFileFromTag(t *testing.T) {
	tag := NewTag(title, artists, album, trackNumber, genre, date, isrc, albumArt())

	flacTag := writeFlacTag(flac.NewFile(256), tag)

	comments := flacTag.VorbisComment.Comments
	if flacTag.VorbisComment.NumberOfComments != 8 {
		t.Errorf("Number of comments should be %d, but was %d", 8, flacTag.VorbisComment.NumberOfComments)
	}
	if comments[0] != title {
		t.Errorf("VorbisComment at %d should be %s but was %s", 0, title, comments[0])
	}
	if comments[1] != artists[0] {
		t.Errorf("VorbisComment at %d should be %s but was %s", 1, artists[0], comments[1])
	}
	if comments[2] != artists[1] {
		t.Errorf("VorbisComment at %d should be %s but was %s", 2, artists[1], comments[2])
	}
	if comments[3] != album {
		t.Errorf("VorbisComment at %d should be %s but was %s", 3, album, comments[3])
	}
	if comments[4] != string(trackNumber) {
		t.Errorf("VorbisComment at %d should be %s but was %s", 4, string(trackNumber), comments[4])
	}
	if comments[5] != genre {
		t.Errorf("VorbisComment at %d should be %s but was %s", 5, genre, comments[5])
	}
	if comments[6] != date {
		t.Errorf("VorbisComment at %d should be %s but was %s", 4, date, comments[6])
	}
	if comments[7] != isrc {
		t.Errorf("VorbisComment at %d should be %s but was %s", 5, isrc, comments[7])
	}
}

func albumArt() []AlbumArt {
	albumArt := make([]AlbumArt, 1)
	imgData := []byte{0x04, 0x08}
	albumArtType := uint32(3)
	mimeType := "image/jpeg"
	width := uint32(50)
	height := uint32(60)
	albumArt[0] = *NewAlbumArt(albumArtType, mimeType, imgData, width, height)

	return albumArt
}
