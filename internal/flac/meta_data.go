package flac

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
