package main

//接口

//类的继承树并无意义，你只需要知道这个类实现了哪些方法，每个方法是啥含义就足够了。
//实现类的时候，只需要关心自己应该提供哪些方法，不用再纠结接口需要拆得多细才合理
//不用为了实现一个接口而导入一个包

//定义File类
type File struct {
	//
}

func (f *File) Read(buf []byte) (n int, err error)                { return 0, nil }
func (f *File) Write(buf []byte) (n int, err error)               { return 0, nil }
func (f *File) Seek(off int64, whence int) (pos int64, err error) { return 0, nil }
func (f *File) Close() error                                      { return nil }

//后定义的IFIle接口
type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}
type IReader interface {
	Read(buf []byte) (n int, err error)
}
type IWriter interface {
	Write(buf []byte) (n int, err error)
}
type ICloser interface {
	Close() error
}

//虽然后定义了这些函数接口，但是依旧可以赋值
func func1() {
	var f1 IFile = new(File)
	var f2 IReader = new(File)
	var f3 IWriter = new(File)
	var f4 ICloser = new(File)
	f1.Close()
	f2.Read(nil)
	f3.Write(nil)
	f4.Close()
}
func main() {
	testInterfaceQuery()
	testTypeQuery()
	testAny()
}
