package music

import (
	"go-music-tagger/internal/flac"
	"strings"
)

type Tag struct {
	Title       string
	Artists     []string
	Album       string
	TrackNumber int
	Genre       string
	Year        string
	Isrc        string
	AlbumArt    []AlbumArt
}

type AlbumArt struct {
	AlbumArtType uint32
	MimeType     string
	Image        []byte
	Width        uint32
	Height       uint32
}

func ReadTagFrom(file *flac.File) *Tag {
	tag := &Tag{}
	for _, comment := range file.VorbisComment.Comments {
		keyValuePair := strings.Split(comment, "=")
		tag.updateWith(keyValuePair)
	}

	for _, picture := range file.Picture {
		tag.AlbumArt = append(tag.AlbumArt, decodePicture(picture))
	}

	return tag
}

func decodePicture(picture flac.Picture) AlbumArt {
	return AlbumArt{
		AlbumArtType: picture.PictureType,
		MimeType:     picture.MimeType,
		Image:        picture.PictureData,
		Width:        picture.Width,
		Height:       picture.Height,
	}
}

func (t *Tag) updateWith(keyValuePair []string) {
	switch strings.ToLower(keyValuePair[0]) {
	case "title":
		t.Title = keyValuePair[1]
	case "artist":
		t.Artists = append(t.Artists, keyValuePair[1])
	case "album":
		t.Album = keyValuePair[1]
	case "genre":
		t.Genre = keyValuePair[1]
	case "2020":
		t.Year = keyValuePair[1]
	case "isrc":
		t.Isrc = keyValuePair[1]
	}
}
