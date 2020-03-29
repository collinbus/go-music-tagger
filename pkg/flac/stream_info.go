package flac

type StreamInfo struct {
	MinimumSampleBlockSize uint16
	MaximumSampleBlockSize uint16
	MinimumFrameSize       uint32
	MaximumFrameSize       uint32
	SampleRate             uint32
	NumberOfChannels       uint8
	BitsPerSample          uint8
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

	streamInfo.MinimumFrameSize = readBigEndianUint32(streamInfoData[4:8], 8)
	streamInfo.MaximumFrameSize = readBigEndianUint32(streamInfoData[7:11], 8)

	streamInfo.SampleRate = readBigEndianUint32(streamInfoData[10:14], 12)
	streamInfo.NumberOfChannels = readBigEndianUint8(streamInfoData[12:14], 4, 9) + 1
	streamInfo.BitsPerSample = readBigEndianUint8(streamInfoData[12:14], 7, 4) + 1
}
