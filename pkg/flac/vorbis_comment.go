package flac

import (
	"bytes"
	"encoding/binary"
)

const sizeOffset = 4

type VorbisComment struct {
	BlockInfo        *BlockInfo
	Vendor           string
	NumberOfComments int
	Comments         []string
}

func NewVorbisComment(blockInfo *BlockInfo) *VorbisComment {
	return &VorbisComment{BlockInfo: blockInfo}
}

func (vc *VorbisComment) Read(data []byte) {
	vendorLength := vc.readVendor(data)
	vc.updateLength(data, vendorLength)

	vc.readComments(data[2*sizeOffset+vendorLength:])
}

func (vc *VorbisComment) updateLength(data []byte, vendorLength uint32) {
	length := binary.LittleEndian.Uint32(data[sizeOffset+vendorLength : vendorLength+2*sizeOffset])
	vc.NumberOfComments = int(length)
}

func (vc *VorbisComment) readVendor(data []byte) uint32 {
	vendorLength := binary.LittleEndian.Uint32(data[:sizeOffset])
	vendorBuffer := bytes.NewBuffer(data[sizeOffset : vendorLength+sizeOffset])
	vc.Vendor = vendorBuffer.String()
	return vendorLength
}

func (vc *VorbisComment) readComments(data []byte) {
	index := 0
	for i := 0; i < vc.NumberOfComments; i++ {
		length := binary.LittleEndian.Uint32(data[index : index+4])
		start := index + sizeOffset
		index += int(length) + sizeOffset
		comment := bytes.NewBuffer(data[start:index]).String()
		vc.Comments = append(vc.Comments, comment)
	}
}
