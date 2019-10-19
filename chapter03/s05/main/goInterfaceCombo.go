package main

//接口组合

// ReadWriter接口将基本的Read和Write方法组合起来
type A interface {
	a()
}
type B interface {
	b()
}

type AB interface {
	A
	B
}
