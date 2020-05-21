package music

import "go-music-tagger/internal/flac"

func writeFlacTag(file *flac.File, tag *Tag) *flac.File {
	file.VorbisComment = createVorbisComments(tag.stringValues())
	return file
}

func createVorbisComments(comments []string) *flac.VorbisComment {
	vorbisComment := flac.NewVorbisComment(flac.NewBlockInfo(0, 0, false))
	vorbisComment.Comments = comments
	vorbisComment.NumberOfComments = len(comments)
	return vorbisComment
}
