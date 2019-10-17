package main

//语言交互性
// 与C语言交互
//Cgo

/*
#include <stdio.h>
*/
import "C"
import "unsafe"

//需要配置GCC
func main() {
	cstr := C.Cstring("H")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
