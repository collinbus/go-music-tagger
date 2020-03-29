package flac

import "encoding/binary"

type StreamInfoReader interface {
	readStreamInfo(streamInfoData []byte)
}

func (streamInfo *StreamInfo) readStreamInfo(streamInfoData []byte) {
	streamInfo.MinimumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[0:2])
	streamInfo.MaximumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[2:4])
}
