package music

import (
	"go-music-tagger/internal/flac"
	"testing"
)

var artists = []string{"an_artist", "a_featuring_artist"}

func TestWriteFlacFileFromTag(t *testing.T) {
	expectedComments := []string{
		"title=" + title,
		"artist=" + artists[0],
		"artist=" + artists[1],
		"album=" + album,
		"tracknumber=1",
		"genre=" + genre,
		"date=" + date,
		"isrc=" + isrc,
	}
	tag := NewTag(title, artists, album, trackNumber, genre, date, isrc, albumArt())

	flacTag := writeFlacTag(flac.NewFile(256), tag)

	comments := flacTag.VorbisComment.Comments
	if flacTag.VorbisComment.NumberOfComments != 8 {
		t.Errorf("Number of comments should be %d, but was %d", 8, flacTag.VorbisComment.NumberOfComments)
	}
	for i, expectedComment := range expectedComments {
		if comments[i] != expectedComment {
			t.Errorf("Comment at %d is %s but should be %s", i, comments[i], expectedComment)
		}
	}
}

func TestWriteFlacFileWithAlbumArtWithTag(t *testing.T) {
	expectedMimeType := "image/jpeg"
	expectedPictureData := []byte{0x04, 0x08}
	expectedWidth := uint32(50)
	expectedHeight := uint32(60)
	tag := NewTag(title, artists, album, trackNumber, genre, date, isrc, albumArt())

	flacTag := writeFlacTag(flac.NewFile(256), tag)
	pictures := flacTag.Pictures

	if pictures[0].PictureType != cover {
		t.Errorf("Pictures type should be %d but was %d", cover, pictures[0].PictureType)
	}
	if pictures[0].MimeType != expectedMimeType {
		t.Errorf("Mime type should be %s but was %s", expectedMimeType, pictures[0].MimeType)
	}
	if pictures[0].PictureData[0] != expectedPictureData[0] {
		t.Errorf("Expected picture data at [0] should be %d but was %d", expectedPictureData, pictures[0].PictureData[0])
	}
	if pictures[0].PictureData[1] != expectedPictureData[1] {
		t.Errorf("Expected picture data at [1] should be %d but was %d", expectedPictureData, pictures[0].PictureData[1])
	}
	if pictures[0].Width != expectedWidth {
		t.Errorf("Width should be %d but was %d", expectedWidth, pictures[0].Width)
	}
	if pictures[0].Height != expectedHeight {
		t.Errorf("Width should be %d but was %d", expectedHeight, pictures[0].Height)
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
