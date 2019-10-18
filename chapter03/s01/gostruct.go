package s01

//结构体
//Go语言放弃了包括继承在内的大量面向对象特性，
//只保留了组合（composition）这个最基础的特性

//矩形类
type Rect struct {
	x, y          float64
	width, height float64
}

//矩形的面积成员方法
func (r *Rect) Area() float64 {
	return r.height * r.width
}

//矩形的周长成员方法
func (r *Rect) Round() float64 {
	return 2 * (r.width + r.height)
}
