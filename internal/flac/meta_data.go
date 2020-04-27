package flac

import "encoding/binary"

type BlockType int

const (
	StreamInfoBlock BlockType = 0
)

type MetaDataReader interface {
	Read(data []byte)
}

type BlockInfo struct {
	startIndex  int
	length      uint32
	isLastBlock bool
}

func NewBlockInfo(startIndex int, length uint32, isLastBlock bool) *BlockInfo {
	return &BlockInfo{startIndex: startIndex, length: length, isLastBlock: isLastBlock}
}

func WriteBlockHeader(isLastBlock bool, blockType BlockType, length uint32) []byte {
	var blockLength = make([]byte, 4)
	blockLength[0] = byte(blockType)
	binary.BigEndian.PutUint32(blockLength, length)
	return blockLength
}
