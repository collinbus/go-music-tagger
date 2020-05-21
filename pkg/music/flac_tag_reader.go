package music

import (
	"go-music-tagger/internal/flac"
	"strings"
)

func readTagFrom(file *flac.File) *Tag {
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
