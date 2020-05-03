package flac

func CreatePadding(length int, audioDataStartIndex int) []byte {
	paddingLength := audioDataStartIndex - length
	padding := make([]byte, paddingLength)
	return padding
}
