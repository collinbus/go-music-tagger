package flac

import "bytes"

type Picture struct {
	BlockInfo   *BlockInfo
	PictureType uint32
	MimeType    string
}

func NewPicture(blockInfo *BlockInfo) *Picture {
	return &Picture{BlockInfo: blockInfo}
}

func (p *Picture) Read(data []byte) {
	var index int
	var end int

	p.PictureType = readBigEndianUint32(data[0:4], 0)

	index = 8
	mimeTypeLength := readBigEndianUint32(data[4:8], 0)
	end = index + int(mimeTypeLength)
	p.MimeType = bytes.NewBuffer(data[index:end]).String()

}
