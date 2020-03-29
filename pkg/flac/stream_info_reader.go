package flac

type StreamInfoReader struct{}

func (*StreamInfoReader) readStreamInfo() *StreamInfo {
	return &StreamInfo{}
}

func NewStreamInfoReader() *StreamInfoReader {
	return &StreamInfoReader{}
}
