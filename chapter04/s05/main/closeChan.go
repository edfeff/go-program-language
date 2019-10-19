package main

import "fmt"

//关闭channel
//关闭channel非常简单，直接使用Go语言内置的close()函数即可
func testCloseChan() {
	c := make(chan int)
	close(c)
	//如何判断一个channel是否已经被关闭
	//在读取的时候使用多重返回值的方式：
	_, ok := <-c
	if ok {
		fmt.Println("没关闭")
	} else {
		fmt.Println("已经关闭了")
	}
}
