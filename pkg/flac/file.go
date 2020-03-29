package flac

type File struct {
	StreamInfo *StreamInfo
	Size       int
}

func NewFile(streamInfo *StreamInfo, size int) *File {
	return &File{StreamInfo: streamInfo, Size: size}
}
