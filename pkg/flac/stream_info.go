package flac

type StreamInfo struct {
	MaximumSampleBlockSize uint16
	MinimumSampleBlockSize uint16
}

func NewStreamInfo() *StreamInfo {
	return &StreamInfo{}
}
