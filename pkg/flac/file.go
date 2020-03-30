package flac

import (
	"errors"
	"io/ioutil"
)

var flacFileIdentifier = []byte{0x66, 0x4C, 0x61, 0x43}

type File struct {
	StreamInfo *StreamInfo
	SeekTable  *SeekTable
	Size       int
}

func NewFile(info *StreamInfo, table *SeekTable, size int) *File {
	return &File{StreamInfo: info, SeekTable: table, Size: size}
}

type BlockInfo struct {
	StartIndex  int
	Length      uint32
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

	info := &StreamInfo{}
	info.BlockInfo = blocks[0]
	fr.metaDataReader.Read(fileBytes[8:42], info)

	seekTable := &SeekTable{}
	seekTable.BlockInfo = blocks[3]
	start := seekTable.BlockInfo.StartIndex
	end := seekTable.BlockInfo.StartIndex + int(seekTable.BlockInfo.Length)
	fr.metaDataReader.Read(fileBytes[start:end], seekTable)

	flacFile := NewFile(info, seekTable, len(fileBytes))
	return flacFile, nil
}

func readMetaDataBlockInfo(data []byte) map[int]BlockInfo {
	const sizeOffset = 4
	var index = sizeOffset
	blocks := make(map[int]BlockInfo)
	for {
		blockId := readBlockId(data[index])
		blockSize := readBigEndianUint32(data[index+1:index+5], 8)
		isLastBlock := readIsLastBlock(data[index])
		blocks[blockId] = BlockInfo{StartIndex: index + sizeOffset, Length: blockSize, isLastBlock: isLastBlock}

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
