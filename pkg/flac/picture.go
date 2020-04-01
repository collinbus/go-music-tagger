package flac

type Picture struct {
	BlockInfo   *BlockInfo
	PictureType uint32
}

func NewPicture(blockInfo *BlockInfo) *Picture {
	return &Picture{BlockInfo: blockInfo}
}

func (p *Picture) Read(data []byte) {
	p.PictureType = readBigEndianUint32(data[0:4], 0)
}
