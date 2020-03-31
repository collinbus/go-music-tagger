package flac

import (
	"bytes"
	"encoding/binary"
)

type VorbisComment struct {
	BlockInfo *BlockInfo
	Vendor    string
}

func NewVorbisComment(blockInfo *BlockInfo) *VorbisComment {
	return &VorbisComment{BlockInfo: blockInfo}
}

func (vc *VorbisComment) Read(data []byte) {
	vendorLength := binary.LittleEndian.Uint32(data[:4])
	buffer := bytes.NewBuffer(data[4:vendorLength])
	vc.Vendor = buffer.String()
}
