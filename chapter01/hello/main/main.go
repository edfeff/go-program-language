//每个Go源代码文件的开头都是一个package声明
// 表示该Go代码所属的包。包是Go语言里最基本的分发单位

//要生成Go可执行程序，必须建立一个名字为main的包，
//并且在该包中包含一个叫main()的函数（该函数是Go可执行程序的执行起点）。
package main

//import语句，用于导入该程序所依赖的包
//不得包含在源代码文件中没有用到的包，否则Go编译器会报编译错误
import "fmt"

//Go语言的main()函数不能带参数，也不能定义返回值。
// 命令行传入的参数在os.Args变量中保存。
// 如果需要支持命令行开关，可使用flag包

//强制左花括号{的放置位置
// 函数名的大小写规则
func main() {
	fmt.Println("Hello go")
}

//所有Go函数（包括在对象编程中会提到的类型成员函数）以关键字func开头
//定义
//func 函数名(参数列表)(返回值列表) {
//	// 函数体
//}
func Compute(v1 int, v2 string, v3 float32) (result1 int64, result2 string, err error) {
	return 0, "", nil
}

//在函数返回时没有被明确赋值的返回值都会被设置为默认值
func Fun(v1 int, v2 string, v3 float32) (result1 int64, result2 string, err error) {
	return
}

//行注释
/*
块注释
*/
//go基本指令
//go version 版本

// 编译并运行程序
// go run XX.go

//编译程序
//go build XX.go

//测试程序
//go test -v

//从根本上说，Go命令行工具只是一个源代码管理工具，或者说是一个前端。
// 真正的Go编译器和链接器被Go命令行工具隐藏在后面，我们可以直接使用它们
//6g和6l是64位版本的Go编译器和链接器，对应的32位版本工具为8g和8l

//
