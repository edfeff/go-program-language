package main

//通道基本语法
//var chanName chan Type

var intChan chan int
var mapChan chan map[string]string
var chanMap map[string]chan bool

//定义一个chan
//ch :=make(chan int)

//写入数据
//ch <-value
//读取数据
//value := <-ch
