package main

//接口赋值
//1 将对象实例赋值给接口
//2 将一个接口赋值给另一个接口

//将对象实例赋值给接口,要求实例实现了该接口的所有方法

type Integer int

//go可以根据值语义函数生成引用语义函数
//如
// func (a *Integer) Less(b Integer) bool {
//	return (*a).Less(b)
//}
//反之则不行，由引用语义函数不能推导出值引用函数
func (a Integer) Less(b Integer) bool {
	return a < b
}

//由于存在 指针引用成员函数，所以对于接口的赋值，只能采用地址赋值
func (a *Integer) Add(b Integer) {
	*a += b
}

//我们定义接口LessAdder
type LessAdder interface {
	Less(b Integer) bool
	Add(b Integer)
}
type Lesser interface {
	Less(b Integer) bool
}

type Adder interface {
	Add(b Integer)
}

func test1() {
	var a Integer = 1
	var b LessAdder = &a
	//不能如下赋值，因为Integer的成员方法中有引用语义函数，所以只能用地址赋值
	//var b LessAdder = a
	b.Less(1)

	//对比 值引用函数 和指针引用函数的赋值不同
	//值引用函数，既可以使用值赋值，也可以采用地址赋值
	var c Lesser = a
	var c2 Lesser = &a
	c.Less(1)
	c2.Less(2)

	var d Adder = &a
	d.Add(1)
}

//所以可以得出结论，当一个类实现了一个接口时，如果这个类中出现了（* X）的成员方法，
//那么将此类赋值到接口时，只能采用地址赋值的方式
