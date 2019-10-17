package main

import (
	"fmt"
	"reflect"
)

// 反射
//  一般用来做对象的序列化(serialization 或称为 Marshal 和 Unmarshal
//标准库中 encoding/json encoding/xml  encoding/gob encoding/binary就是利用反射实现的
//例子
type Bird struct {
	Name string
	Life int
}

func (b *Bird) Fly() {
	fmt.Println("f")
}

func main() {
	sparrow := &Bird{"w", 1}
	s := reflect.ValueOf(sparrow).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d %s % s = %v \n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
	//0 Name string = w
	//1 Life int = 1
}
