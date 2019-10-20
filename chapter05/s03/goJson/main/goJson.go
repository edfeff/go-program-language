package main

import (
	"encoding/json"
	"fmt"
	"log"
)

//Go语言内建对JSON的支持。使用Go语言内置的encoding/json 标准库，
//1 编码为JSON格式

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func EncodeToJson() {
	b := Book{
		Title:       "Go Program",
		Authors:     []string{"wpp"},
		Publisher:   "I.com",
		IsPublished: true,
		Price:       10.4,
	}
	//struct to Json
	bytes, e := json.Marshal(&b)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(string(bytes))
}

// 2 解码JSON数据
func DecodeFromJson() {
	s := `{"Title":"Go Program","Authors":["wpp"],"Publisher":"I.com","IsPublished":true,"Price":10.4}`
	var b Book
	err := json.Unmarshal([]byte(s), &b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.Title)
	fmt.Println(b.Authors)
	fmt.Println(b.Price)

}

//如果JSON中的字段在Go目标类型中不存在，json.Unmarshal()函数在解码过程中会丢弃该字段

// 3 未知格式的JSON

//如果要解码一段未知结构的JSON，只需将这段JSON数据解码输出到一个空接口即可。
//在Go的标准库encoding/json包中，允许使用map[string]interface{}和[]interface{}
//类型的值来分别存放未知结构的JSON对象或数组

func DecodeUnknownJson() {
	s := `{"Title":"Go Program","Authors":["wpp"],"Publisher":"I.com","IsPublished":true,"Price":10.4}`
	var r interface{}
	err := json.Unmarshal([]byte(s), &r)
	//fmt.Println(err, r)
	if err == nil {
		book, ok := r.(map[string]interface{})
		if ok {
			for k, v := range book {
				switch value := v.(type) {
				case string:
					fmt.Println(k, "string", value)
				case int:
					fmt.Println(k, "int", value)
				case bool:
					fmt.Println(k, "bool", value)
				case []interface{}:
					fmt.Println(k, "array", value)
					for i, iv := range value {
						fmt.Println(i, iv)
					}
				default:
					fmt.Println("dont know!")
				}
			}
		}
	}
}
