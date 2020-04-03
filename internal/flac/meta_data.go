package flac

type MetaDataReader interface {
	Read(data []byte)
}
