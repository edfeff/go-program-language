package main

import (
	"errors"
	"fmt"
	"os"
)

//错误处理
//Go语言引入了一个关于错误处理的标准模式，即error接口
//错误接口定义
//type error interface {
//	Error() string
//}

//可以返回错误的函数定义
func Foo(a int) (n int, err error) {
	if a < 0 {
		return 0, errors.New("a < 0 ")
	}
	return a, nil
}

//调用函数的处理模式
func InvokeFoo() {
	n, err := Foo(1)
	if err != nil {
		//错误处理
	} else {
		fmt.Println(n)
	}
}

//定义一个错误结构体
type PathError struct {
	Op   string
	Path string
	Err  error
}

//实现error接口
func (p *PathError) Error() string {
	return p.Op + ":" + p.Path + ";" + p.Err.Error()
}

//use the err
func Stat(name string) (fi os.FileInfo, err error) {
	return nil, &PathError{"op:" + name, "path:" + name, errors.New("path error :" + name)}
}

func CallStat() {
	_, err := Stat("a")
	if err != nil {
		//error type converter
		pathError := err.(*PathError)
		fmt.Println(pathError.Op)
		fmt.Println(pathError.Path)
		fmt.Println(pathError.Err)
	}
}
func main() {
	CallStat()
}
