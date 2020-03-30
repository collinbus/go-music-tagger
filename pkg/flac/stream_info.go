package flac

type StreamInfo struct {
	MinimumSampleBlockSize uint16
	MaximumSampleBlockSize uint16
	MinimumFrameSize       uint32
	MaximumFrameSize       uint32
	SampleRate             uint32
	NumberOfChannels       uint8
	BitsPerSample          uint8
	NumberOfSamples        uint64
	AudioDataMD5Hash       []byte
}

func (streamInfo *StreamInfo) Read(data []byte) {
	streamInfo.MinimumSampleBlockSize = readBigEndianUint16(data[0:2])
	streamInfo.MaximumSampleBlockSize = readBigEndianUint16(data[2:4])

	streamInfo.MinimumFrameSize = readBigEndianUint32(data[4:8], 8)
	streamInfo.MaximumFrameSize = readBigEndianUint32(data[7:11], 8)

	streamInfo.SampleRate = readBigEndianUint32(data[10:14], 12)
	streamInfo.NumberOfChannels = readBigEndianUint8(data[12:14], 4, 9) + 1
	streamInfo.BitsPerSample = readBigEndianUint8(data[12:14], 7, 4) + 1

	streamInfo.NumberOfSamples = readBigEndianUint64(data[13:21], 4, 24)
	streamInfo.AudioDataMD5Hash = data[18:34]
}
