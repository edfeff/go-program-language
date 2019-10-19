package main

import "fmt"

//缓冲机制
//
func testChanBuffer() {
	//调用make()时将缓冲区大小作为第二个参数传入即可
	//，即使没有读取方，写入方也可以一直往channel里写入，在缓冲区被填完之前都不会阻塞
	c := make(chan int, 1024)

	//从带缓冲的channel中读取数据可以使用与常规非缓冲channel完全一致的方法，
	// 但我们也可以使用range关键来实现更为简便的循环读取
	for i := range c {
		fmt.Println(i)
	}

}
