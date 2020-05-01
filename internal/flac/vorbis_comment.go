package flac

import (
	"bytes"
	"encoding/binary"
)

const sizeOffset = 4
const vendor = "go-music-tagger v0.0.1"

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
	length := readLittleEndianUint32(data[sizeOffset+vendorLength : vendorLength+2*sizeOffset])
	vc.NumberOfComments = int(length)
}

func (vc *VorbisComment) readVendor(data []byte) uint32 {
	vendorLength := readLittleEndianUint32(data[:sizeOffset])
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

func (vc *VorbisComment) WriteVorbisComments() []byte {
	buffer := make([]byte, 0)
	vendorBytes := bytes.NewBufferString(vendor).Bytes()

	vendorLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(vendorLength, uint32(len(vendorBytes)))

	commentsLength := make([]byte, 4)
	binary.LittleEndian.PutUint32(commentsLength, uint32(vc.NumberOfComments))

	buffer = append(buffer, vendorLength...)
	buffer = append(buffer, vendorBytes...)
	buffer = append(buffer, commentsLength...)

	for _, comment := range vc.Comments {
		commentBytes := bytes.NewBufferString(comment).Bytes()
		commentLength := make([]byte, 4)
		binary.LittleEndian.PutUint32(commentLength, uint32(len(commentBytes)))

		buffer = append(buffer, commentLength...)
		buffer = append(buffer, commentBytes...)
	}

	blockSize := len(buffer)
	blockHeader := WriteBlockHeader(false, VorbisCommentBlock, uint32(blockSize))
	return append(blockHeader, buffer...)
}
