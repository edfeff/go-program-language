package main

import "fmt"

//流程控制
//条件语句 if / else / else if
//选择语句 switch / case / select
//循环 for / range
//跳转 goto
//中断,跳出 break continue fallthrough

func main() {
	//conditionFunc()
	//selectionFunc()
	gotoFunc()
}

//条件语句
func conditionFunc() {
	a := 80
	if a > 100 {
		fmt.Println("A")
	} else if a > 80 {
		fmt.Println("B")
	} else if a > 60 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
}

//选择语句
func selectionFunc() {
	s := "A"
	switch s {
	case "A":
		fmt.Println("100-80")
	case "B":
		fmt.Println("80-70")
	case "C":
		fmt.Println("70-60")
	default:
		fmt.Println("<60")
	}
	//可省略表达式,但case的值可以计算成布尔
	num := 100
	switch {
	case num > 100:
		fmt.Println(">100")
	case num > 80:
		fmt.Println(">80")
	case num > 60:
		fmt.Println(">60")
	default:
		fmt.Println("<60")
	}
}

//循环
func loopFunc() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	//while true
	sum = 0
	for {
		if sum > 100 {
			break
		}
		sum++
	}

	//break and continue
	for i := 0; i < 10; i++ {
		if i == 2 {
			break
		}
	}
	for i := 0; i < 10; i++ {
		if i == 2 {
			continue
		}
	}
}

//goto
func gotoFunc() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i < 10 {
		goto HERE
	}
}
