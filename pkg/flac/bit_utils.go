package flac

import (
	"encoding/binary"
)

func readBigEndianUint8(bytes []byte, startOffset int, endOffset int) uint8 {
	numberBeforeShift := binary.BigEndian.Uint16(bytes)
	numberWithoutEndOffset := numberBeforeShift << startOffset
	number := numberWithoutEndOffset >> (startOffset + endOffset)
	return uint8(number)
}

func readBigEndianUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

func readBigEndianUint32(bytes []byte, offset int) uint32 {
	numberBeforeShift := binary.BigEndian.Uint32(bytes)
	number := numberBeforeShift >> offset
	return number
}
