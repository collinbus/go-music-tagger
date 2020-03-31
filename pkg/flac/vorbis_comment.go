package flac

type VorbisComment struct {
	BlockInfo *BlockInfo
}

func NewVorbisComment(blockInfo *BlockInfo) *VorbisComment {
	return &VorbisComment{BlockInfo: blockInfo}
}

func (vc *VorbisComment) Read(data []byte) {

}
