package main

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter01/calcproj/calc"
	"github.com/wpp/go-program-language/chapter01/calcproj/simplemath"
	"os"
	"strconv" //字符串转换
)

//工程结构
func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		calc.Usage()
		return
	}
	//for i, a := range args {
	//	fmt.Println(i, a)
	//}
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("calc add <int> <int>")
			return
		}
		v1, e1 := strconv.Atoi(args[2])
		v2, e2 := strconv.Atoi(args[3])
		if e1 != nil || e2 != nil {
			fmt.Println("calc add <int> <int>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("R:", ret)
	case "sqrt":
		if len(args) != 3 {
			fmt.Println("calc sqrt <int>")
			return
		}
		v1, e1 := strconv.Atoi(args[2])
		if e1 != nil {
			fmt.Println("calc sqrt <int>")
			return
		}
		ret := simplemath.Sqrt(v1)
		fmt.Println("R:", ret)
	default:
		calc.Usage()
	}

}
