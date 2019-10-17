package main

import "fmt"

//枚举 一般用常量定义
const (
	Sunday       = iota //0
	Monday              //1
	Tuesday             //2
	Wednesday           //3
	Thursday            //4
	Friday              //5
	Saturday            //6
	numberOfDays        //7不公开
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)
}
