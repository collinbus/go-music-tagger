package flac

import (
	"log"
	"os"
)

func WriteFile(source File, target string) *os.File {
	var buffer = make([]byte, 0)

	flacHeader := writeFlacHeader()
	streamInfo := source.StreamInfo.WriteStreamInfoBlock()
	seekPoints := source.SeekTable.WriteSeekPoints()
	vorbisComments := source.VorbisComment.WriteVorbisComments()
	pictureData := make([]byte, 0)

	for _, picture := range source.Picture {
		pictureBytes := picture.WritePicture()
		pictureData = append(pictureData, pictureBytes...)
	}

	padding := CreatePadding(len(buffer), source.audioDataStart)

	buffer = append(buffer, flacHeader...)
	buffer = append(buffer, streamInfo...)
	buffer = append(buffer, seekPoints...)
	buffer = append(buffer, vorbisComments...)
	buffer = append(buffer, pictureData...)
	buffer = append(buffer, padding...)
	buffer = append(buffer, source.AudioData...)

	newFile, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}
	_, _ = newFile.Write(buffer)
	return newFile
}

func writeFlacHeader() []byte {
	return []byte{0x66, 0x4C, 0x61, 0x43}
}
