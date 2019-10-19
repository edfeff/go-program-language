package main

import "runtime"

//并行求和
//在Go语言升级到默认支持多CPU的某个版本之前，我们可以先通过设置环境变量
//GOMAXPROCS的值来控制使用多少个CPU核心。具体操作方法是通过直接设置环境变量
//GOMAXPROCS的值，或者在代码中启动goroutine之前先调用以下这个语句以设置使用16个CPU
//核心：
//runtime.GOMAXPROCS(16)
type Vector []float64

func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
}

func (v Vector) Op(f float64) float64 {
	return f
}

const NCPU = 4

func (v Vector) DoAll(u Vector) {
	//设置调度器个数
	runtime.GOMAXPROCS(NCPU)
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		//每个goroutine只计算一部分
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	for i := 0; i < NCPU; i++ {
		<-c //获取到一个数据，表示一个CPU计算完成了
	}
}
