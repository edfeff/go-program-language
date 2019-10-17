package main

import "fmt"

//基础类型
// 布尔
var b1 bool = true
var b2 bool = false

//符号整性
var i int = 1     // 平台相关
var i8 int8 = 1   // [-128,127]
var i16 int16 = 1 //[-32768,32767]
var i32 int32 = 1 //[-2147483648 , 2147483647]
var i64 int64 = 1 //[-9 223 372 036 854 775 808 ,9 223 372 036 854 775 807]

//无符号整性
var u uint = 1     // 平台相关
var u8 uint8 = 1   //[0,255]
var u16 uint16 = 1 //[0,65535]
var u32 uint32 = 1 //[0 ~ 4 294 967 295]
var u64 uint64 = 1 //[0 ~ 18 446 744 073 709 551 615]

//浮点
//自动推导的类型默认为float64
var f32 float32 = 1.0
var f64 float64 = 1.0

//复数
var c64 complex64 = 1 + 1i
var c128 complex128 = 1 + 1i

//字符串
var s string = ""

//字符
var r rune = 1

//错误类型
var e error = nil

func main() {
	sf()
}

//复合类型
//指针
var p *int = &i

//数组
var arr [2]int

//切片
var splice []int

//映射
var m map[string]string

//通道
var c chan int

//结构体
var st struct{}

//接口
var interface1 interface{}

//
func bf() {
	bb1 := (1 == 2)
	fmt.Println(bb1)
	//不能强转
	//var bb2 bool
	//bb2 = 1 //错误
	//bb2 = bool(1) //错误
}

//字符串遍历的两种方式
func sf() {
	ss1 := "Hello,世界"
	//用字节数组的方式遍历,取出的类型是byte
	for i := 0; i < len(ss1); i++ {
		fmt.Println(i, ss1[i])
	}
	//0 72
	//1 101
	//2 108
	//3 108
	//4 111
	//5 44
	//6 228
	//7 184
	//8 150
	//9 231
	//10 149
	//11 140

	//用字符的方式遍历,取出的类型是rune(按照 Unicode 方式的取值)
	for i, x := range ss1 {
		fmt.Println(i, x)
	}
	//0 72
	//1 101
	//2 108
	//3 108
	//4 111
	//5 44
	//6 19990
	//9 30028
}

//字符
//支持2种字符类型,
// 一个是byte,实际是uint8,用于表示单字节的值,
// 一个是rune,表示单个Unicode字符
func cf() {

}
