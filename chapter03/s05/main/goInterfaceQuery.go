package main

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter03/s05/two"
)

//接口查询
//有办法让上面的Writer接口转换为two.IStream接口么？有。那就是我们即将讨论的接口查询语法

func testInterfaceQuery() {
	var f1 Writer = new(File)
	//在Go语言中，你可以询问接口它指向的对象是否是某个类型
	//即f1的实例是不是完全实现了two.IStream
	if f2, ok := f1.(two.IStream); ok {
		fmt.Println("接口查询成功", f2)
	}
}
