package main

//让出时间片
//可以在每个goroutine中控制何时主动出让时间片
//runtime 包中的Gosched()函数实现。
