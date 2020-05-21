package music

import "go-music-tagger/internal/flac"

func writeFlacTag(file *flac.File, tag *Tag) *flac.File {
	file.VorbisComment = createVorbisComments(tag.stringValues())
	file.Pictures = createPictures(tag.AlbumArt)
	return file
}

func createPictures(art []AlbumArt) []flac.Picture {
	pictures := make([]flac.Picture, len(art))
	for i, albumArt := range art {
		picture := flac.NewPicture(nil)
		picture.PictureType = albumArt.AlbumArtType
		picture.MimeType = albumArt.MimeType
		picture.PictureData = albumArt.Image
		picture.Width = albumArt.Width
		picture.Height = albumArt.Height

		pictures[i] = *picture
	}
	return pictures
}

func createVorbisComments(comments []string) *flac.VorbisComment {
	vorbisComment := flac.NewVorbisComment(flac.NewBlockInfo(0, 0, false))
	vorbisComment.Comments = comments
	vorbisComment.NumberOfComments = len(comments)
	return vorbisComment
}
