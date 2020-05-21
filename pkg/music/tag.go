package music

import "strings"

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

func (tag *Tag) stringValues() []string {
	tagValues := make([]string, 0)
	tagValues = append(tagValues, tag.Title)
	tagValues = append(tagValues, tag.Artists...)
	tagValues = append(tagValues, tag.Album)
	tagValues = append(tagValues, string(tag.TrackNumber))
	tagValues = append(tagValues, tag.Genre)
	tagValues = append(tagValues, tag.Year)
	tagValues = append(tagValues, tag.Isrc)
	return tagValues
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
	case "date":
		t.Year = keyValuePair[1]
	case "isrc":
		t.Isrc = keyValuePair[1]
	}
}

func NewTag(title string,
	artists []string,
	album string,
	trackNumber int,
	genre string,
	year string,
	isrc string,
	albumArt []AlbumArt,
) *Tag {
	return &Tag{
		Title:       title,
		Artists:     artists,
		Album:       album,
		TrackNumber: trackNumber,
		Genre:       genre,
		Year:        year,
		Isrc:        isrc,
		AlbumArt:    albumArt,
	}
}

type AlbumArt struct {
	AlbumArtType uint32
	MimeType     string
	Image        []byte
	Width        uint32
	Height       uint32
}

func NewAlbumArt(
	albumArtType uint32,
	mimeType string,
	image []byte,
	width uint32,
	height uint32,
) *AlbumArt {
	return &AlbumArt{
		AlbumArtType: albumArtType,
		MimeType:     mimeType,
		Image:        image,
		Width:        width,
		Height:       height,
	}
}
