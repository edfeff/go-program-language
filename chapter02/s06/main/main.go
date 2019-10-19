package main

import (
	"errors"
	"fmt"
)

//函数
//多返回值函数
//首字母大写,为public
//首字母小写,为private
//小写字母开头的函数只在本包内可见，大写字母开头的函数才能被其他包使用

func Add(a int, b int) (result int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("a or b is less than 0")
		return
	}
	return a + b, nil
}

//不定参数,指定参数类型
func myPrint(args ...int) {
	for _, v := range args {
		fmt.Println(v)
	}
}

//不定参数,任意类型
func mySuperPrint(args ...interface{}) {
	for _, v := range args {
		fmt.Println(v)
	}
}

// 不定参数的类型
func findType(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case int64:
			fmt.Println("int64")
		default:
			fmt.Println("dont know")
		}
	}
}

func main() {
	//不定参数的传递
	//1 普通多参数
	myPrint(1, 2, 3)
	mySuperPrint(4, 5, 6)

	//2 切片参数
	s := []int{2, 2, 2}
	myPrint(s[0:]...)

	//
	findType("1", 1, 111, 1.1)
	//匿名函数
	noNameFunction()
	closureTest()
}

//匿名函数与闭包
//在Go里面，函数可以像普通变量一样被传递或使用
func noNameFunction() {
	//1. 匿名函数直接赋值给变量
	f1 := func(a int, b int) int {
		return a + b
	}
	//通过变量调用
	i := f1(1, 2)
	fmt.Println(i)
	//2. 直接调用
	func() {
		fmt.Println("direct invoke")
	}()
	//3 接受参数
	b := func(a, b int) bool {
		return a > b
	}(1, 2)
	fmt.Println(b)
}
