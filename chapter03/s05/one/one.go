package one

//one 和 two都具有相同声明的接口，可以互相赋值
type ReadWriter interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
}
