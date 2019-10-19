package main

import "fmt"

//select
//select有比较多的限制，其中最大的一条限制就是每个case语句里必须是一个IO操作
//select不像switch，后面并不带判断条件，而是直接去查看case语句。每个case语句都必须是一个面向channel的操作

func testSelect() {
	c1 := make(chan int)
	c2 := make(chan int)
	//第一个case试图从chan1读取一个数据并直接忽略读到的数据，
	//而第二个case则是试图向chan2中写入一个整型数1，
	//如果这两者都没有成功，则到达default语句。
	select {
	case v := <-c1:
		fmt.Println(v)
	case c2 <- 1:
		fmt.Println()
	default:
		break
	}
}

//死循环select
func testSelect2() {
	ch := make(chan int, 1)
	for {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		fmt.Println(i)
	}
}
