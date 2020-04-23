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
	streamInfo := writeStreamInfoBlock(source.StreamInfo)

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
func writeStreamInfoBlock(info *StreamInfo) []byte {
	var minimumBlockSize = make([]byte, 2)
	var maximumBlockSize = make([]byte, 2)
	var minimumFrameSize = make([]byte, 4)
	var maximumFrameSize = make([]byte, 4)
	var otherInfoBytes = make([]byte, 8)
	var md5Signature = info.AudioDataMD5Hash
	var streamInfo = make([]byte, 0)

	binary.BigEndian.PutUint16(minimumBlockSize, info.MinimumSampleBlockSize)
	binary.BigEndian.PutUint16(maximumBlockSize, info.MaximumSampleBlockSize)
	binary.BigEndian.PutUint32(minimumFrameSize, info.MinimumFrameSize)
	binary.BigEndian.PutUint32(maximumFrameSize, info.MaximumFrameSize)

	var otherInfo uint64
	otherInfo = uint64(info.SampleRate)
	otherInfo = otherInfo << 3
	otherInfo += uint64(info.NumberOfChannels - 1)
	otherInfo = otherInfo << 5
	otherInfo += uint64(info.BitsPerSample - 1)
	otherInfo = otherInfo << 36
	otherInfo += info.NumberOfSamples

	binary.BigEndian.PutUint64(otherInfoBytes, otherInfo)

	streamInfo = append(streamInfo, minimumBlockSize...)
	streamInfo = append(streamInfo, maximumBlockSize...)
	streamInfo = append(streamInfo, minimumFrameSize[1:4]...)
	streamInfo = append(streamInfo, maximumFrameSize[1:4]...)
	streamInfo = append(streamInfo, otherInfoBytes...)
	streamInfo = append(streamInfo, md5Signature...)
	return streamInfo
}

//noinspection GoNilness
func writeBlockHeader(info *StreamInfo) []byte {
	var blockLength = make([]byte, 4)
	binary.BigEndian.PutUint32(blockLength, info.BlockInfo.length)
	return blockLength
}
