package main

import "github.com/wpp/go-program-language/chapter05/s01/dial"

//网络编程

//Socket编程
//Go语言标准库里提供的net包，支持基于IP层、TCP/UDP层及更高层面（如HTTP、FTP、SMTP）的网络操作，
//其中用于IP层的称为RawSocket。
//
func main() {
	//dial.TestDialICMP()
	//dial.TestTCP()
	dial.TestTCP2()
}
