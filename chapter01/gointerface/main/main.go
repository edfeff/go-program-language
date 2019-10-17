package main

import "fmt"

//接口与实现松耦合,不用强制声明实现和继承关系

type Bird struct {
	Name string
	Age  int
}

func (b *Bird) Fly() {
	fmt.Println(b.Name, "fly")
}
func main() {
	a := Bird{"wpp", 1}
	a.Fly()

	var b IFly = &Bird{
		Name: "bb",
		Age:  0,
	}
	b.Fly()

}

type IFly interface {
	Fly()
}
