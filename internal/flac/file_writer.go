package flac

import (
	"encoding/binary"
	"log"
	"os"
)

type FileWriter interface {
	WriteFile(target string) (bool, error)
}

type FileWriterService struct {
	source *File
}

func NewFileWriterService(source *File) *FileWriterService {
	return &FileWriterService{source: source}
}

func (fs FileWriterService) WriteFile(target string) (bool, error) {
	newFile, err := os.Create(target)
	if err != nil {
		log.Fatal(err)
	}

	fs.writeDataTo(newFile)
	return true, nil
}

func (fs FileWriterService) writeDataTo(file *os.File) {
	fs.writeFlacHeader(file)
	fs.writeStreamInfoBlock(fs.source.StreamInfo, file)
}

func (fs FileWriterService) writeFlacHeader(file *os.File) {
	_, _ = file.Write([]byte{0x66, 0x4C, 0x61, 0x43})
}

//noinspection GoNilness
func (fs FileWriterService) writeStreamInfoBlock(info *StreamInfo, file *os.File) {
	writeBlockInfo(file, info)
	var minimumBlockSize = make([]byte, 2)
	var maximumBlockSize = make([]byte, 2)
	var minimumFrameSize = make([]byte, 4)
	var maximumFrameSize = make([]byte, 4)
	var otherInfoBytes = make([]byte, 8)
	var md5Signature = info.AudioDataMD5Hash

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

	_, _ = file.Write(minimumBlockSize)
	_, _ = file.Write(maximumBlockSize)
	_, _ = file.Write(minimumFrameSize[1:4])
	_, _ = file.Write(maximumFrameSize[1:4])
	_, _ = file.Write(otherInfoBytes)
	_, _ = file.Write(md5Signature)
}

//noinspection GoNilness
func writeBlockInfo(file *os.File, info *StreamInfo) {
	_, _ = file.Write([]byte{0})
	var length = make([]byte, 4)
	binary.BigEndian.PutUint32(length, info.BlockInfo.length)
	_, _ = file.Write((length)[1:4])
}
