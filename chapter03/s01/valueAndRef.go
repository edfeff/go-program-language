package s01

import "fmt"

//值语义与引用语义
func TestValueAndRef() {
	test1()
	test2()
	testIntegerRef()
}

//值复制
func test1() {
	var a = [3]int{1, 2, 3}
	var b = a
	b[1]++
	fmt.Println(a, b)
	//[1 2 3] [1 3 3]
}

//引用复制
func test2() {
	var a = [3]int{1, 2, 3}
	var b = &a
	b[1]++
	fmt.Println(a, *b)
	//[1 3 3] [1 3 3]
}

//切片的模拟定义,所以即使是切片赋值的话,也是值复制
type slice struct {
	first *interface{}
	len   int
	cap   int
}

//map定义类比
//map就是一个指针字典
type K string

type V string

type MAP_K_V struct {
}
type MyMap struct {
	impl *MAP_K_V
}

//基于指针设计一个引用类型
type IntegerRef struct {
	impl *int
}

func testIntegerRef() {
	var a int = 10
	var pa IntegerRef = IntegerRef{&a}
	*pa.impl = 999
	fmt.Println(a, *pa.impl)
}

//所以go的类型都是值语义的,引用语义是通过封装指针实现的
