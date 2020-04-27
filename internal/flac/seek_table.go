package flac

import "encoding/binary"

type SeekPoint struct {
	FirstSampleNumber    uint64
	Offset               uint64
	NumberOfTargetFrames uint16
}

type SeekTable struct {
	BlockInfo          *BlockInfo
	NumberOfSeekPoints int
	SeekPoints         []SeekPoint
}

func NewSeekTable(blockInfo *BlockInfo) *SeekTable {
	return &SeekTable{BlockInfo: blockInfo}
}

func (s *SeekTable) Read(data []byte) {
	numberOfSeekPoints := s.BlockInfo.length / 18
	s.NumberOfSeekPoints = int(numberOfSeekPoints)

	for i := 0; i < s.NumberOfSeekPoints; i++ {
		index := i * 18
		seekPoint := SeekPoint{
			FirstSampleNumber:    readBigEndianUint64(data[index:index+8], 0, 0),
			Offset:               readBigEndianUint64(data[index+8:index+16], 0, 0),
			NumberOfTargetFrames: readBigEndianUint16(data[index+16 : index+18]),
		}
		s.SeekPoints = append(s.SeekPoints, seekPoint)
	}
}

func (s *SeekTable) WriteSeekPoints() []byte {
	seekTableBlockLength := s.BlockInfo.length
	buffer := WriteBlockHeader(false, SeekTableBlock, seekTableBlockLength)

	for _, sp := range s.SeekPoints {
		var firstSample = make([]byte, 8)
		var offset = make([]byte, 8)
		var numberOfSamples = make([]byte, 2)

		binary.BigEndian.PutUint64(firstSample, sp.FirstSampleNumber)
		binary.BigEndian.PutUint64(offset, sp.Offset)
		binary.BigEndian.PutUint16(numberOfSamples, sp.NumberOfTargetFrames)

		buffer = append(buffer, firstSample...)
		buffer = append(buffer, offset...)
		buffer = append(buffer, numberOfSamples...)
	}

	return buffer
}
