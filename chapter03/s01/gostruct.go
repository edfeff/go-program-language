package s01

//结构体
//Go语言放弃了包括继承在内的大量面向对象特性，
//只保留了组合（composition）这个最基础的特性

//矩形类
type Rect struct {
	X, Y          float64
	Width, Height float64
}

//矩形的面积成员方法
func (r *Rect) Area() float64 {
	return r.Height * r.Width
}

//矩形的周长成员方法
func (r *Rect) Round() float64 {
	return 2 * (r.Width + r.Height)
}

//构造函数
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}
