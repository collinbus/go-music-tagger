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
