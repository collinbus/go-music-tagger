package flac

type SeekPoint struct {
	FirstSampleNumber    uint64
	Offset               uint64
	NumberOfTargetFrames uint16
}

type SeekTable struct {
	BlockInfo          BlockInfo
	NumberOfSeekPoints int
	SeekPoints         []SeekPoint
}

func (s *SeekTable) Read(data []byte) {
	numberOfSeekPoints := s.BlockInfo.length / 18
	s.NumberOfSeekPoints = int(numberOfSeekPoints)

	seekPoint := SeekPoint{
		FirstSampleNumber:    readBigEndianUint64(data[0:8], 0, 0),
		Offset:               readBigEndianUint64(data[8:16], 0, 0),
		NumberOfTargetFrames: readBigEndianUint16(data[16:18]),
	}
	s.SeekPoints = append(s.SeekPoints, seekPoint)
}
