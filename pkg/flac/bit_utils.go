package flac

import "encoding/binary"

func readBigEndianUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

func readBigEndianUint32(bytes []byte, offset int) uint32 {
	numberBeforeShift := binary.BigEndian.Uint32(bytes)
	number := numberBeforeShift >> offset
	return number
}
