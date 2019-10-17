package s04

import "fmt"

type PersonInfo struct {
	Name    string
	Address string
}

func RunMap() {
	//创建map
	//var personDB map[string]PersonInfo
	personDB := make(map[string]PersonInfo)

	//插入数据
	personDB["1"] = PersonInfo{"wpp", "home"}
	personDB["2"] = PersonInfo{"ccc", "Come"}

	//查找获取
	person, ok := personDB["1"]
	if ok {
		fmt.Printf("Yes,found in DB!\nName is %s,Address is %s\n", person.Name, person.Address)
	} else {
		fmt.Println("can't find 1 in DB")
	}
	//删除
	delete(personDB, "1")
	person, ok = personDB["1"]
	if ok {
		fmt.Printf("Yes,found in DB!\nName is %s,Address is %s\n", person.Name, person.Address)
	} else {
		fmt.Println("Sorry! Can't find 1 in DB")
	}
}
