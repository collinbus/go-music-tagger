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
	var index int
	var end int

	p.PictureType = readBigEndianUint32(data[0:4], 0)

	index = 8
	mimeTypeLength := readBigEndianUint32(data[4:8], 0)
	end = index + int(mimeTypeLength)
	p.MimeType = bytes.NewBuffer(data[index:end]).String()

	index = end
	end = index + 4
	descriptionLength := readBigEndianUint32(data[index:end], 0)
	index += 4
	end = index + int(descriptionLength)
	p.Description = bytes.NewBuffer(data[index:end]).String()

	index = end
	end += 4
	p.Width = readBigEndianUint32(data[index:end], 0)

	index = end
	end += 4
	p.Height = readBigEndianUint32(data[index:end], 0)

	index = end
	end += 4
	p.ColorDepth = readBigEndianUint32(data[index:end], 0)

	index = end
	end += 4
	p.IndexedColorPictures = readBigEndianUint32(data[index:end], 0)

	index += 8
	p.PictureData = data[index:]
}
