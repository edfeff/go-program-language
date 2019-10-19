package main

import "fmt"

//类型查询
//在Go语言中，还可以更加直截了当地询问接口指向的对象实例的类型
func testTypeQuery() {
	var v1 interface{} = 1
	//类型查询
	switch v := v1.(type) {
	case int:
		fmt.Println(v, "int")
	case float32:
		fmt.Println(v, "float32")
	default:
		fmt.Println(v, "i dont know")
	}
}
