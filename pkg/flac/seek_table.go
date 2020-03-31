package flac

type SeekTable struct {
	BlockInfo          BlockInfo
	NumberOfSeekPoints int
}

func (s *SeekTable) Read(data []byte) {
	numberOfSeekPoints := s.BlockInfo.length / 18
	s.NumberOfSeekPoints = int(numberOfSeekPoints)
}
