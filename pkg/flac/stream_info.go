package flac

type StreamInfo struct {
	MaximumSampleBlockSize uint16
	MinimumSampleBlockSize uint16
}

func NewStreamInfo(maximumSampleBlockSize uint16, minimumSampleBlockSize uint16) *StreamInfo {
	return &StreamInfo{
		MaximumSampleBlockSize: maximumSampleBlockSize,
		MinimumSampleBlockSize: minimumSampleBlockSize,
	}
}
