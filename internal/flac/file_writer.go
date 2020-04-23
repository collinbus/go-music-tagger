package flac

import (
	"log"
	"os"
)

func WriteFile(source File, target string) *os.File {
	var buffer = make([]byte, 0)

	flacHeader := writeFlacHeader()
	blockHeader := source.StreamInfo.WriteBlockHeader()
	streamInfo := source.StreamInfo.WriteStreamInfoBlock()

	buffer = append(buffer, flacHeader...)
	buffer = append(buffer, blockHeader...)
	buffer = append(buffer, streamInfo...)

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
