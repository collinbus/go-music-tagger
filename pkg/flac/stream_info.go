package flac

type StreamInfo struct {
	MinimumSampleBlockSize uint16
	MaximumSampleBlockSize uint16
	MinimumFrameSize       uint32
	MaximumFrameSize       uint32
}

func NewStreamInfo() *StreamInfo {
	return &StreamInfo{}
}
