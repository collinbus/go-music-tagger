package flac

type StreamInfoReader struct{}

type StreamInfo struct {
}

func (*StreamInfoReader) readStreamInfo() *StreamInfo {
	return &StreamInfo{}
}

func NewStreamInfoReader() *StreamInfoReader {
	return &StreamInfoReader{}
}
