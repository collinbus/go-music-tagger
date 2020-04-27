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

func readBigEndianUint64(bytes []byte, startOffset int, endOffset int) uint64 {
	numberBeforeShift := binary.BigEndian.Uint64(bytes)
	numberWithoutEndOffset := numberBeforeShift << startOffset
	number := numberWithoutEndOffset >> (startOffset + endOffset)
	return uint64(number)
}

func readLittleEndianUint32(data []byte) uint32 {
	return binary.LittleEndian.Uint32(data)
}

func setBit(b byte, pos uint) byte {
	b |= 1 << pos
	return b
}
