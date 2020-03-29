package flac

import "encoding/binary"

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
	streamInfo.MinimumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[0:2])
	streamInfo.MaximumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[2:4])

	streamInfo.MinimumFrameSize = readBigEndian24BitAsUint32(streamInfoData[4:7])
	streamInfo.MaximumFrameSize = readBigEndian24BitAsUint32(streamInfoData[7:10])
}
