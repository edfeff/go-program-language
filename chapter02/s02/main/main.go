package main

//常量定义
const PI float64 = 3.1415926
const zero = 0
const (
	k   int64 = 1024
	eof int   = -1
)
const u, v, w = 1, "A", 1.00
const mask = 1 << 3

//预定义常量
// iota比较特殊，可以被认为是一个可被编译器修改的常量，
// 在每一个const关键字出现时被重置为0，然后在下一个const出现之前，每出现一次iota，其所代表的数字会自动增1
const (
	c1 = iota //0
	c2        //1
	c3        //2
)
const (
	a = 1 << iota //1
	b             //2
	c             //4
	d             //8
)
const (
	U         = iota * 42 //0
	V float64 = iota * 42 //42.0
	W         = iota * 42 //84
)
const x = iota //0
const y = iota //0
const z = iota //0

func main() {
}
