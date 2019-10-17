package main

import "fmt"

//goroutine是一种比线程更加轻盈、更省资源的协程

//并行计算的例子
func sum(values []int, resultChan chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	resultChan <- sum
}
func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	//分成两组计算
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	s1, s2 := <-resultChan, <-resultChan
	fmt.Println(s1, s2, s1+s2)
}
