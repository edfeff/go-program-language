package main

import "fmt"

//可见性
//大写字母开头--外部包（package）可见
//Go语言中符号的可访问性是包一级的而不是类型一级的
//尽管area()是Rect的内部方法，但同一个包中的其他类型也都可以访问到它
type Rect struct {
	X, Y          float64
	Width, Height float64
}

func (r *Rect) area() float64 {
	return r.Width * r.Height
}

func main() {
	hello()
	b := body{}
	fmt.Println(b)
}
