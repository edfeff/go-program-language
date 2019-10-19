package main

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter03/s05/one"
	"github.com/wpp/go-program-language/chapter03/s05/two"
)

//接口赋值给接口
//在Go语言中，只要两个接口拥有相同的方法列表（次序不同不要紧），那么它们就是等同的，可以相互赋值
//one 和 two都具有相同声明的接口，可以互相赋值
//在Go语言中，这两个接口实际上并无区别，因为：
// 任何实现了one.ReadWriter接口的类，均实现了two.IStream；
// 任何one.ReadWriter接口对象可赋值给two.IStream，反之亦然；
// 在任何地方使用one.ReadWriter接口与使用two.IStream并无差异。

func compareInterface() {
	var f1 one.ReadWriter = new(File)
	var f2 two.IStream = new(File)
	var f3 two.IStream = f1

	fmt.Println(f1, f2, f3)
}

//接口赋值并不要求两个接口必须等价。如果接口A的方法列表是接口B的方法列表的子集，
//那么接口B可以赋值给接口A。
type Writer interface {
	Write(buf []byte) (n int, err error)
}

func innerInterface() {
	var f1 one.ReadWriter = new(File)
	var f2 two.IStream = new(File)
	//Writer方法是one.ReadWriter 或two.IStream方法的子集
	var f3 Writer = f1
	var f4 Writer = f2
	fmt.Println(f1, f2, f3, f4)
}
