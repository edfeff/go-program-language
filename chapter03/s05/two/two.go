package two

type IStream interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}
