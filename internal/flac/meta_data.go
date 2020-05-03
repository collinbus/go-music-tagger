package flac

import "encoding/binary"

type BlockType uint32

const (
	StreamInfoBlock    BlockType = 0
	PaddingBlock                 = 1
	SeekTableBlock               = 3
	VorbisCommentBlock           = 4
	PictureBlock                 = 6
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
	binary.BigEndian.PutUint32(blockLength, length)

	blockLength[0] = byte(blockType)

	if isLastBlock {
		blockLength[0] = setBit(blockLength[0], 7)
	}

	return blockLength
}
