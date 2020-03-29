package flac

import "encoding/binary"

func readBigEndian24BitAsUint32(bytes []byte) uint32 {
	var number = []byte{0x00}
	number = append(number, bytes...)
	return binary.BigEndian.Uint32(number)
}
