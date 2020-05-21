package music

import (
	"go-music-tagger/internal/flac"
	"strconv"
)

func writeFlacTag(file *flac.File, tag *Tag) *flac.File {
	file.VorbisComment = createVorbisComments(tag)
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

func createVorbisComments(tag *Tag) *flac.VorbisComment {
	comments := vorbisStringValues(tag)
	vorbisComment := flac.NewVorbisComment(flac.NewBlockInfo(0, 0, false))
	vorbisComment.Comments = comments
	vorbisComment.NumberOfComments = len(comments)
	return vorbisComment
}

func vorbisStringValues(tag *Tag) []string {
	tagValues := make([]string, 0)
	tagValues = append(tagValues, writeVariable("title", tag.Title))
	for _, artist := range tag.Artists {
		tagValues = append(tagValues, writeVariable("artist", artist))
	}
	tagValues = append(tagValues, writeVariable("album", tag.Album))
	tagValues = append(tagValues, writeVariable("tracknumber", strconv.FormatInt(int64(tag.TrackNumber), 10)))
	tagValues = append(tagValues, writeVariable("genre", tag.Genre))
	tagValues = append(tagValues, writeVariable("date", tag.Year))
	tagValues = append(tagValues, writeVariable("isrc", tag.Isrc))
	return tagValues
}

func writeVariable(name string, value string) string {
	return name + "=" + value
}
