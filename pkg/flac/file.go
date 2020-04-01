package flac

import (
	"errors"
	"io/ioutil"
)

var flacFileIdentifier = []byte{0x66, 0x4C, 0x61, 0x43}

type File struct {
	StreamInfo    *StreamInfo
	SeekTable     *SeekTable
	VorbisComment *VorbisComment
	Picture       []Picture
	Size          int
}

func NewFile(size int) *File {
	return &File{Size: size, Picture: []Picture{}}
}

type BlockInfo struct {
	startIndex  int
	length      uint32
	isLastBlock bool
}

type FileReader struct {
	metaDataReader MetaDataReader
}

func (fr *FileReader) ReadFile(path string) (*File, error) {
	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	if !isFlacFile(fileBytes) {
		return nil, errors.New("file at " + path + " is not a flac file")
	}

	blocks := readMetaDataBlockInfo(fileBytes)
	flacFile := NewFile(len(fileBytes))

	for blockId, blockInfo := range blocks {
		flacFile.readMetaData(fileBytes, blockId, blockInfo)
	}

	return flacFile, nil
}

func (f *File) readMetaData(data []byte, blockId int, info *BlockInfo) {
	start := info.startIndex
	end := start + int(info.length)

	switch blockId {
	case 0:
		f.StreamInfo = NewStreamInfo(info)
		f.StreamInfo.Read(data[start:end])
	case 3:
		f.SeekTable = NewSeekTable(info)
		f.SeekTable.Read(data[start:end])
	case 4:
		f.VorbisComment = NewVorbisComment(info)
		f.VorbisComment.Read(data[start:end])
	case 6:
		picture := NewPicture(info)
		picture.Read(data)
		f.Picture = append(f.Picture, *picture)
	}
}

func readMetaDataBlockInfo(data []byte) map[int]*BlockInfo {
	const sizeOffset = 4
	var index = sizeOffset
	blocks := make(map[int]*BlockInfo)
	for {
		blockId := readBlockId(data[index])
		blockSize := readBigEndianUint32(data[index+1:index+5], 8)
		isLastBlock := readIsLastBlock(data[index])
		blocks[blockId] = &BlockInfo{startIndex: index + sizeOffset, length: blockSize, isLastBlock: isLastBlock}

		if isLastBlock {
			break
		}

		index += int(blockSize) + sizeOffset
	}
	return blocks
}

func readIsLastBlock(b byte) bool {
	isLastBlock := b >> 7
	return isLastBlock == 1
}

func readBlockId(b byte) int {
	blockId := b << 1 >> 1
	return int(blockId)
}

func isFlacFile(data []byte) bool {
	if len(data) < 4 {
		return false
	}

	for i := range flacFileIdentifier {
		if data[i] != flacFileIdentifier[i] {
			return false
		}
	}
	return true
}

func NewFileReader(metaDataReader MetaDataReader) *FileReader {
	return &FileReader{
		metaDataReader: metaDataReader,
	}
}
