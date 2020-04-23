package flac

import (
	"encoding/binary"
	"log"
	"os"
)

func WriteFile(source File, target string) *os.File {
	var buffer = make([]byte, 0)

	flacHeader := writeFlacHeader()
	blockHeader := writeBlockHeader(source.StreamInfo)
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

//noinspection GoNilness
func writeBlockHeader(info *StreamInfo) []byte {
	var blockLength = make([]byte, 4)
	binary.BigEndian.PutUint32(blockLength, info.BlockInfo.length)
	return blockLength
}
