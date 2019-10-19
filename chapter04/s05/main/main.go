package main

import (
	"fmt"
)

//channel
//channel是Go语言在语言级别提供的goroutine间的通信方式。
// 我们可以使用channel在两个或多个goroutine之间传递消息。
// channel是进程内的通信方式，因此通过channel传递对象的过程和调用函数时的参数传递行为比较一致，比如也可以传递指针等。

//channel是类型相关的。也就是说，一个channel只能传递一种类型的值，这个类型需要在声明channel时指定。

func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting")
}
func testCount() {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	for _, ch := range chs {
		//直接取出丢弃
		<-ch
	}
}
func main() {
	//testSelect2()
	testWriteAndReadBySingleDirectionChannel()

}
