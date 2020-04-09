package flac

import (
	"errors"
	"io/ioutil"
)

type FileReaderService struct {
	metaDataReader MetaDataReader
}

type FileReader interface {
	ReadFile(path string) (*File, error)
}

func (fr *FileReaderService) ReadFile(path string) (*File, error) {
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

	flacFile.AudioData = fileBytes[flacFile.audioDataStart:]

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

	if info.isLastBlock {
		f.audioDataStart = info.startIndex + int(info.length)
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

func NewFileService(metaDataReader MetaDataReader) *FileReaderService {
	return &FileReaderService{
		metaDataReader: metaDataReader,
	}
}
