package music

import (
	"go-music-tagger/internal/flac"
	"testing"
)

const title = "a_title"
const album = "an_album"
const artist = "an_artist"
const genre = "a_genre"
const date = "2020"
const isrc = "an_isrc"

func TestReadTagFromFile(t *testing.T) {
	tagReader := NewTagReader(&MockFileReader{})

	tag := tagReader.ReadTagFrom("a_file_path")

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
	if tag.Isrc != isrc {
		t.Errorf("Expected %s, but was %s", isrc, tag.Isrc)
	}
}

type MockFileReader struct{}

func (m *MockFileReader) ReadFile(path string) (*flac.File, error) {
	comments := []string{"title=" + title, "ARTIST=" + artist, "Album=" + album,
		"genre=" + genre, "date=" + date, "ISRC=" + isrc}
	vorbisComment := &flac.VorbisComment{Comments: comments}
	return &flac.File{VorbisComment: vorbisComment}, nil
}
