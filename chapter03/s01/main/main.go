package main

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter03/s01"
)

//类型系统

//给内置类型添加方法

//定义类型别名
type Integer int

func (a Integer) Less(b Integer) bool {
	return a < b
}

//使用面向过程的方式重写上面的代码
//go 本质也是这样做的,外面用语法包裹起来了
//将this指针用语法糖封装起来
func Integer_Less(a Integer, b Integer) bool {
	return a < b
}

//使用指针去修改原有的对象
func (a *Integer) Add(b Integer) {
	*a += b
}

//传递值,不能改变原来的值
//是因为Go语言和C语言一样，类型都是基于值传递的。要想修改变量的值，只能传递指针
func (a Integer) NAdd(b Integer) {
	a += b
}

//测试内置的方法
func testInteger() {
	var a Integer = 1
	if a.Less(2) {
		fmt.Println("less")
	}
	a.Add(1)
	fmt.Println(a) //2

	a.NAdd(1)
	fmt.Println(a) //2
}
func main() {
	//testInteger()
	s01.TestValueAndRef()
}
