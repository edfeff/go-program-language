package main

import (
	"fmt"
	"time"
)

//defer 关键字
//保证了函数退出时一定被调用
//defer匿名函数
//defer的顺序,defer是栈结构//defer 关键字
////保证了函数退出时一定被调用
////defer匿名函数
////defer的顺序,defer是栈结构
type Job struct {
	name string
}

func (j *Job) Start() {
	fmt.Println(j.name + "Start...")

}

func (j *Job) Close() {
	fmt.Println(j.name + "Close...")
}
func UserJob() {
	j := Job{"wpp"}
	defer j.Close()
	j.Start()
	time.Sleep(time.Second * 2)
}

//defer匿名函数
func noNameDefer() {
	fmt.Println("11111")
	defer func() {
		fmt.Println("END")
	}()
	fmt.Println("222222")
}

//defer的顺序,defer是栈结构
func deferOrder() {
	defer func() {
		fmt.Println(1)
	}()
	defer func() {
		fmt.Println(2)
	}()
	defer func() {
		fmt.Println(3)
	}()
	defer func() {
		fmt.Println(4)
	}()
	//4
	//3
	//2
	//1
}
