package music

import (
	"go-music-tagger/internal/flac"
	"reflect"
	"testing"
)

const title = "a_title"
const album = "an_album"
const artist = "an_artist"
const genre = "a_genre"
const date = "2020"
const isrc = "an_isrc"

const mimeType = "a_mime_type"
const height = 500
const width = 300
const cover = 3

var albumImage = []byte{0x00, 0x01}

func TestReadTagFromFile(t *testing.T) {
	tag := readTagFrom(aFlacFile())

	if tag.Title != title {
		t.Errorf("Expected %s, but was %s", title, tag.Title)
	}
	if tag.Album != album {
		t.Errorf("Expected %s, but was %s", album, tag.Album)
	}
	if tag.Artists[0] != artist {
		t.Errorf("Expected %s, but was %s", artist, tag.Artists[0])
	}
	if tag.Genre != genre {
		t.Errorf("Expected %s, but was %s", genre, tag.Genre)
	}
	if tag.Year != date {
		t.Errorf("Expected %s, but was %s", date, tag.Year)
	}
	if tag.Isrc != isrc {
		t.Errorf("Expected %s, but was %s", isrc, tag.Isrc)
	}
}

func TestReadAlbumArtFromFile(t *testing.T) {
	tag := readTagFrom(aFlacFile())

	albumArt := tag.AlbumArt[0]
	if albumArt.MimeType != mimeType {
		t.Errorf("Expected %s, but was %s", mimeType, albumArt.MimeType)
	}
	if albumArt.Height != height {
		t.Errorf("Expected %d, but was %d", height, albumArt.Height)
	}
	if albumArt.Width != width {
		t.Errorf("Expected %d, but was %d", width, albumArt.Width)
	}
	if albumArt.AlbumArtType != cover {
		t.Errorf("Expected %d, but was %d", cover, albumArt.AlbumArtType)
	}
	if !reflect.DeepEqual(albumArt.Image, albumImage) {
		t.Error("The image's binary data is incorrect")
	}
}

func aFlacFile() *flac.File {
	comments := []string{"title=" + title, "ARTIST=" + artist, "Album=" + album,
		"genre=" + genre, "date=" + date, "ISRC=" + isrc}
	vorbisComment := &flac.VorbisComment{Comments: comments}
	picture := &flac.Picture{
		PictureType: cover,
		MimeType:    mimeType,
		PictureData: albumImage,
		Width:       width,
		Height:      height,
	}
	return &flac.File{VorbisComment: vorbisComment, Picture: []flac.Picture{*picture}}
}
