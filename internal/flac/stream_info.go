package flac

import "encoding/binary"

type StreamInfo struct {
	BlockInfo              *BlockInfo
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

func NewStreamInfo(blockInfo *BlockInfo) *StreamInfo {
	return &StreamInfo{BlockInfo: blockInfo}
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

func (streamInfo *StreamInfo) WriteBlockHeader() []byte {
	var blockLength = make([]byte, 4)
	binary.BigEndian.PutUint32(blockLength, streamInfo.BlockInfo.length)
	return blockLength
}

func (streamInfo *StreamInfo) WriteStreamInfoBlock() []byte {
	var minimumBlockSize = make([]byte, 2)
	var maximumBlockSize = make([]byte, 2)
	var minimumFrameSize = make([]byte, 4)
	var maximumFrameSize = make([]byte, 4)
	var otherInfoBytes = make([]byte, 8)
	var md5Signature = streamInfo.AudioDataMD5Hash
	var streamInfoBytes = make([]byte, 0)

	binary.BigEndian.PutUint16(minimumBlockSize, streamInfo.MinimumSampleBlockSize)
	binary.BigEndian.PutUint16(maximumBlockSize, streamInfo.MaximumSampleBlockSize)
	binary.BigEndian.PutUint32(minimumFrameSize, streamInfo.MinimumFrameSize)
	binary.BigEndian.PutUint32(maximumFrameSize, streamInfo.MaximumFrameSize)

	var otherInfo uint64
	otherInfo = uint64(streamInfo.SampleRate)
	otherInfo = otherInfo << 3
	otherInfo += uint64(streamInfo.NumberOfChannels - 1)
	otherInfo = otherInfo << 5
	otherInfo += uint64(streamInfo.BitsPerSample - 1)
	otherInfo = otherInfo << 36
	otherInfo += streamInfo.NumberOfSamples

	binary.BigEndian.PutUint64(otherInfoBytes, otherInfo)

	streamInfoBytes = append(streamInfoBytes, minimumBlockSize...)
	streamInfoBytes = append(streamInfoBytes, maximumBlockSize...)
	streamInfoBytes = append(streamInfoBytes, minimumFrameSize[1:4]...)
	streamInfoBytes = append(streamInfoBytes, maximumFrameSize[1:4]...)
	streamInfoBytes = append(streamInfoBytes, otherInfoBytes...)
	streamInfoBytes = append(streamInfoBytes, md5Signature...)
	return streamInfoBytes
}
