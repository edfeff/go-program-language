package main

import (
	"fmt"
	"time"
)

//单向chan
//我们在将一个channel变量传递到一个函数时，可以通过将其指定为单向channel变量，从
//而限制该函数中可以对此channel的操作，比如只能往这个channel写，或者只能从这个
//channel读。

var ch1 chan int
var ch2 chan<- float64 //单向channel ，只能用于写入float64
var ch3 <-chan int     // 单向channel，只能用于读取int数据

//对chan进行类型转换，实现反向的channel
func testSingleChannel() {
	ch4 := make(chan int)
	//ch5 := <-(ch4) 简化写法
	ch5 := <-chan int(ch4) //ch5 就是一个单向的读取channel
	ch6 := chan<- int(ch4) //ch6 就是一个单向的写入channel

	fmt.Println(ch4, ch5, ch6)
}
func testWriteAndReadBySingleDirectionChannel() {
	baseChannel := make(chan int)
	//rchan := <-chan int(baseChannel) //读通道
	//wchan := chan<- int(baseChannel) //写通道

	go writeOnlyChannel(baseChannel)
	go readOnlyChannel(baseChannel)
	time.Sleep(1e9)
}

//单向channel的用法
//单向读取
func readOnlyChannel(ch <-chan int) {
	//不能发送数据，除非进行类型转换
	//ch <- 999
	for value := range ch {
		fmt.Println(value)
	}
}

//单向写入
func writeOnlyChannel(ch chan<- int) {
	for i := 1; i < 10; i++ {
		ch <- i
	}
}
