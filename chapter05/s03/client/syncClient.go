package client

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter05/s03/server"
	"log"
	"net/rpc"
)

const serverAddress = "localhost:12345"

//同步调用
func ArithMultiplySync(a, b int) int {
	//创建客户端
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	args := server.Args{a, b}

	var reply int

	//同步调用
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d * %d = %d\n", args.A, args.B, reply)
	return reply
}
func ArithDivideSync(a, b int) *server.Quotient {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	//同步调用
	args := server.Args{a, b}
	var reply server.Quotient
	err = client.Call("Arith.Divide", args, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Arith: %d / %d = %d ,%d\n", args.A, args.B, reply.Quo, reply.Rem)
	return &reply
}

//调用 加减
func AddSubAdd(a, b int) int {
	client, e := rpc.DialHTTP("tcp", serverAddress)
	if e != nil {
		log.Fatal(e)
	}
	args := server.Args{a, b}
	var reply int
	e = client.Call("AddSub.Add", &args, &reply)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(reply)
	return reply
}
func AddSubSub(a, b int) int {
	client, e := rpc.DialHTTP("tcp", serverAddress)
	if e != nil {
		log.Fatal(e)
	}
	args := server.Args{a, b}
	var reply int
	e = client.Call("AddSub.Sub", &args, &reply)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(reply)
	return reply
}
