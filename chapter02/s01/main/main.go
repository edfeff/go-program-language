package main

import "fmt"

//变量声明
var v1 int
var v2 string
var v3 [10]int  //数组
var v4 []int    //切片
var v5 struct { //结构体
	f int
}
var v6 *int            // 指针
var v7 map[int]string  // 映射
var v8 func(a int) int //函数

//变量组合声明
var (
	vv1 int
	vv2 string
)

//初始化
var vvv1 int = 10
var vvv2 = 10

//vvv3 := 10 //必须在函数内声明

//赋值
func f1() {
	var x1 int
	x1 = 10
	fmt.Println(x1)
}

//交换赋值
func f2() {
	i, j := 1, 2
	fmt.Println(i, j)
	i, j = j, i
	fmt.Println(i, j)
}

//匿名变量,不需要使用的
func f3() {
	_, name := f4()
	fmt.Println(name)
}
func f4() (int, string) {
	return 0, ""
}

func main() {
	f2()
}
