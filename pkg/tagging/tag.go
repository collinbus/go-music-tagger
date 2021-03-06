package tagging

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
