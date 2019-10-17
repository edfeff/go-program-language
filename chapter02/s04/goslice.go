package s04

import "fmt"

//切片
//数组切片的数据结构可以抽象为以下3个变量：
//1. 一个指向原生数组的指针；
//2. 数组切片中的元素个数；
//3. 数组切片已分配的存储空间
func RunSlice() {
	//createSplice()
	//traversalArr()
	//dynamicOp()
	contentCopy()
}

//创建数组
func createSplice() {
	//1 基于数组
	var arr = [3]int{1, 2, 3}
	//，Go语言支持用myArray[first:last)这样的方式来基于数组生成一 个数组切片
	var s1 []int = arr[:]   //全部元素
	var s2 []int = arr[1:]  //从第二个元素开始到结束
	var s3 []int = arr[:1]  //从头到第二个
	var s4 []int = arr[1:1] //空白
	var s5 []int = arr[1:2]

	fmt.Println(s1) //[1 2 3]
	fmt.Println(s2) //[2 3]
	fmt.Println(s3) //[1]
	fmt.Println(s4) //[]
	fmt.Println(s5) //[2]

	//2 直接创建
	ms1 := make([]int, 5)           //长度5 的切片
	ms2 := make([]int, 5, 10)       //长度5,容量10 的切片
	ms3 := []int{1, 2, 3}           //直接初始化
	fmt.Println(len(ms1), cap(ms1)) //5 5
	fmt.Println(len(ms2), cap(ms2)) //5 10
	fmt.Println(len(ms3), cap(ms3)) //3 3
}

//遍历切片
func traversalArr() {
	s1 := []int{1, 2, 3}
	//方式1
	for i := 0; i < len(s1); i++ {
		fmt.Println(i, s1[i])
	}
	//方式2
	for index, value := range s1 {
		fmt.Println(index, value)
	}
}

//动态增删
func dynamicOp() {
	s1 := make([]int, 5, 10)
	s2 := []int{8, 9, 10}
	fmt.Println(len(s1), cap(s1))
	//添加元素
	s1 = append(s1, 1, 2, 3) //5 10
	//添加别的对象使用 ... 语法
	s1 = append(s1, s2...)
	fmt.Println(len(s1), cap(s1)) //11 20
}

//内容复制
//如果加入的两个数组切片不一样大，
// 就会按其中较小的那个数组切片的元素个数进行复制
func contentCopy() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3}
	copy(s1, s2)
	fmt.Println(s1, s2) //[5 4 3 4 5] [5 4 3]
	contentCopyLargeToSmall()
}

func contentCopyLargeToSmall() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{5, 4, 3}
	copy(s2, s1)
	fmt.Println(s1, s2) //[1 2 3 4 5] [1 2 3]
}
