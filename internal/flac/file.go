package flac

var flacFileIdentifier = []byte{0x66, 0x4C, 0x61, 0x43}

type File struct {
	StreamInfo     *StreamInfo
	SeekTable      *SeekTable
	VorbisComment  *VorbisComment
	Picture        []Picture
	Size           int
	audioDataStart int
	AudioData      []byte
}

func NewFile(size int) *File {
	return &File{Size: size, Picture: []Picture{}}
}
