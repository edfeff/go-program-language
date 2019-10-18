package main

import "log"

//匿名组合
//Go语言也提供了继承，但是采用了组合的文法，所以我们将其称为匿名组合
type Base struct {
	Name string
}

func (b *Base) BaseFoo() {}
func (b *Base) BaseBar() {}

//匿名继承
type Foo struct {
	Base
	Address string
}

//新方法
func (f *Foo) FooFoo() {}

//重载
func (f *Foo) BaseFoo() {}

//指针派生
type Bar struct {
	*Base
	Age int
}

//官方指针继承示例
type Worker struct {
	Name        string
	*log.Logger //组合一个日志指针
}

func (w *Worker) Start() {
	w.Printf("%s start...", w.Name)
	//...
	w.Printf("%s end...", w.Name)
}

//接口组合中的名字冲突问题
type X struct {
	Name string
}
type Y struct {
	X
	Name string // X的Name被隐藏了
}

func main() {

}
