package flac

import "encoding/binary"

type StreamInfoReader struct{}

func (*StreamInfoReader) readStreamInfo(streamInfoData []byte) *StreamInfo {
	minimumSampleBlockSize := binary.BigEndian.Uint16(streamInfoData[0:2])
	maximumSampleBlockSize := binary.BigEndian.Uint16(streamInfoData[2:4])

	return NewStreamInfo(maximumSampleBlockSize, minimumSampleBlockSize)
}

func NewStreamInfoReader() *StreamInfoReader {
	return &StreamInfoReader{}
}
