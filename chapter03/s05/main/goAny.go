package main

//Any类型
//Go语言中任何对象实例都满足空接口interface{}，所以interface{}看起来像是可以指向任何对象的Any类型，

var v1 interface{} = 1     // 将int类型赋值给interface{}
var v2 interface{} = "abc" // 将string类型赋值给interface{}
var v3 interface{} = &v2   // 将*interface{}类型赋值给interface{}
var v4 interface{} = struct{ X int }{1}
var v5 interface{} = &struct{ X int }{1}

//当函数可以接受任意的对象实例时，我们会将其声明为interface{}，
func any(i interface{}) {

}
func testAny() {
	any(v1)
	any(v2)
	any(v3)
	any(v4)
	any(v5)
}
