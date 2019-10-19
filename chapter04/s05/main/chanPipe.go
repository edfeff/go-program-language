package main

//chan 传递
//在Go语言中channel本身也是一个原生类型，与map之类的类型地位一样，因此channel本身在定义后也可以通过channel来传递

//我们可以使用这个特性来实现*nix上非常常见的管道（pipe）特性。管道也是使用非常广泛
//的一种设计模式，比如在处理数据时，我们可以采用管道设计，这样可以比较容易以插件的方式
//增加数据的处理流程

//传递int类型的管道
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

//流式处理
func handle(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}
