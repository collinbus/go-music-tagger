package flac

import (
	"bytes"
	"encoding/binary"
)

type Picture struct {
	BlockInfo            *BlockInfo
	PictureType          uint32
	MimeType             string
	Description          string
	Width                uint32
	Height               uint32
	ColorDepth           uint32
	IndexedColorPictures uint32
	PictureData          []byte
}

func NewPicture(blockInfo *BlockInfo) *Picture {
	return &Picture{BlockInfo: blockInfo}
}

func (p *Picture) Read(data []byte) {
	var index = p.BlockInfo.startIndex
	var end = index + 4

	p.PictureType = readBigEndianUint32(data[index:end], 0)

	index, end = p.updateMimeType(index, end, data)

	index, end = p.updateDescription(index, end, data)

	index, end = p.updateIndexForNextBlock(index, end)
	p.Width = readBigEndianUint32(data[index:end], 0)

	index, end = p.updateIndexForNextBlock(index, end)
	p.Height = readBigEndianUint32(data[index:end], 0)

	index, end = p.updateIndexForNextBlock(index, end)
	p.ColorDepth = readBigEndianUint32(data[index:end], 0)

	index, end = p.updateIndexForNextBlock(index, end)
	p.IndexedColorPictures = readBigEndianUint32(data[index:end], 0)

	index, end = p.updateIndexForNextBlock(index, end)
	pictureDataLength := int(readBigEndianUint32(data[index:end], 0))

	index = end
	end += pictureDataLength
	p.PictureData = data[index:end]
}

func (p *Picture) updateMimeType(index int, end int, data []byte) (int, int) {
	index, end = p.updateIndexForNextBlock(index, end)
	mimeTypeLength := readBigEndianUint32(data[index:end], 0)
	index = end
	end = index + int(mimeTypeLength)
	p.MimeType = bytes.NewBuffer(data[index:end]).String()
	return index, end
}

func (p *Picture) updateDescription(index int, end int, data []byte) (int, int) {
	index, end = p.updateIndexForNextBlock(index, end)
	descriptionLength := readBigEndianUint32(data[index:end], 0)
	index += 4
	end = index + int(descriptionLength)
	p.Description = bytes.NewBuffer(data[index:end]).String()
	return index, end
}

func (p *Picture) updateIndexForNextBlock(index int, end int) (int, int) {
	index = end
	end = index + 4
	return index, end
}

func (p *Picture) WritePicture() []byte {
	buffer := make([]byte, 0)

	pictureType := make([]byte, 4)
	binary.BigEndian.PutUint32(pictureType, p.PictureType)
	mimeTypeLength := make([]byte, 4)
	binary.BigEndian.PutUint32(mimeTypeLength, uint32(len(p.MimeType)))
	mimeType := bytes.NewBufferString(p.MimeType).Bytes()
	descriptionLength := make([]byte, 4)
	binary.BigEndian.PutUint32(descriptionLength, uint32(len(p.Description)))
	description := bytes.NewBufferString(p.Description).Bytes()
	width := make([]byte, 4)
	binary.BigEndian.PutUint32(width, p.Width)
	height := make([]byte, 4)
	binary.BigEndian.PutUint32(height, p.Height)
	colorDepth := make([]byte, 4)
	binary.BigEndian.PutUint32(colorDepth, p.ColorDepth)
	indexedColorPictures := make([]byte, 4)
	binary.BigEndian.PutUint32(indexedColorPictures, p.IndexedColorPictures)
	lengthPictureData := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthPictureData, uint32(len(p.PictureData)))

	buffer = append(buffer, pictureType...)
	buffer = append(buffer, mimeTypeLength...)
	buffer = append(buffer, mimeType...)
	buffer = append(buffer, descriptionLength...)
	buffer = append(buffer, description...)
	buffer = append(buffer, width...)
	buffer = append(buffer, height...)
	buffer = append(buffer, colorDepth...)
	buffer = append(buffer, indexedColorPictures...)
	buffer = append(buffer, lengthPictureData...)
	buffer = append(buffer, p.PictureData...)

	header := WriteBlockHeader(false, PictureBlock, uint32(len(buffer)))
	return append(header, buffer...)
}
