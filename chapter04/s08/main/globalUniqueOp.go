package main

import (
	"fmt"
	"sync"
	"time"
)

//全局唯一性操作
//对于从全局的角度只需要运行一次的代码，比如全局初始化操作，Go语言提供了一个Once 类型来保证全局的唯一性操作

var a string
var once sync.Once

func setup() {
	a = "hello a"
	fmt.Println("init a")
}
func doPrint() {
	//once的Do()方法可以保证在全局范围内只调用指定的函数一次（这里指
	//setup()函数），而且所有其他goroutine在调用到此语句时，将会先被阻塞，直至全局唯一的
	//once.Do()调用结束后才继续
	once.Do(setup)
	fmt.Println(a)
}
func twoPrint() {
	go doPrint()
	go doPrint()
	time.Sleep(1e9)
}
