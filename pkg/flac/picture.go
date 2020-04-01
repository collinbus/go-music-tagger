package flac

type Picture struct {
	BlockInfo *BlockInfo
}

func NewPicture(blockInfo *BlockInfo) *Picture {
	return &Picture{BlockInfo: blockInfo}
}

func (p *Picture) Read(data []byte) {

}
