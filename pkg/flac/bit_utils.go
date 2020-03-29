package flac

import "encoding/binary"

func readBigEndianUint16(bytes []byte) uint16 {
	return binary.BigEndian.Uint16(bytes)
}

func readBigEndianUint32(bytes []byte, offset int) uint32 {
	var number []byte
	bytesToBeAdded := offset / 8
	for i := 0; i < bytesToBeAdded; i++ {
		number = append(number, 0x00)
	}

	number = append(number, bytes...)
	return binary.BigEndian.Uint32(number)
}
