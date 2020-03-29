package flac

import "encoding/binary"

type StreamInfoReader interface {
	readStreamInfo(streamInfoData []byte)
}

func (streamInfo *StreamInfo) readStreamInfo(streamInfoData []byte) {
	streamInfo.MinimumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[0:2])
	streamInfo.MaximumSampleBlockSize = binary.BigEndian.Uint16(streamInfoData[2:4])

	streamInfo.MinimumFrameSize = readBigEndian24BitAsUint32(streamInfoData[4:7])
	streamInfo.MaximumFrameSize = readBigEndian24BitAsUint32(streamInfoData[7:10])
}

func readBigEndian24BitAsUint32(bytes []byte) uint32 {
	var number = []byte{0x00}
	number = append(number, bytes...)
	return binary.BigEndian.Uint32(number)
}
