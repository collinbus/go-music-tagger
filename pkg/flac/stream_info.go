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

type StreamInfoReader interface {
	readStreamInfo(streamInfoData []byte)
}

func (streamInfo *StreamInfo) readStreamInfo(streamInfoData []byte) {
	streamInfo.MinimumSampleBlockSize = readBigEndianUint16(streamInfoData[0:2])
	streamInfo.MaximumSampleBlockSize = readBigEndianUint16(streamInfoData[2:4])

	streamInfo.MinimumFrameSize = readBigEndianUint32(streamInfoData[4:7], 8)
	streamInfo.MaximumFrameSize = readBigEndianUint32(streamInfoData[7:10], 8)
}
