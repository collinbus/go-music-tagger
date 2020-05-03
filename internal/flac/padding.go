package flac

func CreatePadding(length int, audioDataStartIndex int) []byte {
	paddingLength := audioDataStartIndex - length
	header := WriteBlockHeader(true, PaddingBlock, uint32(paddingLength))
	padding := make([]byte, paddingLength)
	return append(header, padding...)
}
