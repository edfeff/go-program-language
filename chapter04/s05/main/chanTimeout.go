package main

import "time"

//超时机制
//Go语言没有提供直接的超时处理机制，但我们可以利用select机制。虽然select机制不是
//专为超时而设计的，却能很方便地解决超时问题。
//因为select的特点是只要其中一个case已经完成，程序就会继续往下执行，而不会考虑其他case的情况。

//基于此特性，我们来为channel实现超时机制
func testTimeOut(ch chan interface{}, chHandler func(interface{})) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) //1s 超时
		timeout <- true
	}()
	//如果1s没有从ch中取出数据，就从timeout中取出数据，结束整个程序，避免阻塞等待
	select {
	case v := <-ch:
		//从ch中读取数据,交给传入的函数
		chHandler(v)
	case <-timeout:
		//一直没有从ch中读取到数据，但从timeout中读取到了数据
	}
}
