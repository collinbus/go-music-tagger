package music

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
