package main

import (
	"fmt"
	"log"
)

//go中引入了2个内置函数panic()和recover()以报告和处理运行时错误和程序中的错误场景
//func panic(interface{})
//func recover(interface{})

//错误处理流程
//当在一个函数执行过程中调用panic()函数时，正常的函数执行流程将立即终止，但函数中
//之前使用defer关键字延迟执行的语句将正常展开执行，之后该函数将返回到调用函数，并导致
//逐层向上执行panic流程，直至所属的goroutine中所有正在执行的函数被终止。错误信息将被报
//告，包括在调用panic()函数时传入的参数，这个过程称为错误处理流程。

//recover()函数用于终止错误处理流程。一般情况下，recover()应该在一个使用defer
//关键字的函数中执行以有效截取错误处理流程。如果没有在发生异常的goroutine中明确调用恢复
//过程（使用recover关键字），会导致该goroutine所属的进程打印异常信息后直接退出。

func testPanic() {
	f1()
}
func f1() {
	defer func() {
		//捕获异常，并进行处理
		if r := recover(); r != nil {
			log.Println("f1 caught panic:", r)
		}
	}()
	f2()
}

func f2() {
	defer func() {
		fmt.Println("f2 defer")
	}()
	fmt.Println("before panic")
	panic("f2 error")
	fmt.Println("after panic")

}
