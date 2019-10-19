package main

//同步
//即使成功地用channel来作为通信手段，还是避免不了多个goroutine之间共享数据的问题，
// Go 语言的设计者虽然对channel有极高的期望，但也提供了妥善的资源锁方案。

func main() {
	twoPrint()
}
