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

	buffer = append(buffer, flacHeader...)
	buffer = append(buffer, streamInfo...)
	buffer = append(buffer, seekPoints...)

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
