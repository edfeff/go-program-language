package main

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter03/s01"
)

//初始化
//在Go语言中，未进行显式初始化的变量都会被初始化为该类型的零值，
//对象的创建通常交由一个全局的创建函数来完成，以NewXXX来命名，表示“构造函数”
func main() {
	r1 := new(s01.Rect)
	r2 := &s01.Rect{}
	r3 := &s01.Rect{0, 0, 100, 100}
	r4 := &s01.Rect{Width: 100, Height: 100}
	r5 := s01.NewRect(1, 1, 1, 1)
	fmt.Println(r1, r2, r3, r4, r5)
}
