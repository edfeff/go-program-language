package main

import "fmt"

//goroutine
func Add(x, y int) {
	fmt.Println(x + y)
}
func goAdd() {
	for i := 0; i < 10; i++ {
		go Add(1, 2)
	}
}
func main() {
	testCount()
}
