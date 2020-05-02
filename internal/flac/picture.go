package flac

import "bytes"

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

	index += 8
	p.PictureData = data[index:]
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
