package client

import (
	"fmt"
	"github.com/wpp/go-program-language/chapter05/s03/server"
	"log"
	"net/rpc"
)

//异步调用rpc
func ArithDivideASync(a, b int) {
	result := new(server.Quotient)
	args := server.Args{a, b}
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal(err)
	}
	//异步调用
	divCall := client.Go("Arith.Divide", args, &result, nil)

	//从通道中取出答复值
	reply := <-divCall.Done

	fmt.Println(reply)
	if reply.Error != nil {
		log.Fatal(reply.Error)
	} else {
		r := reply.Reply
		fmt.Println("r:", r)
	}
	fmt.Println("*result:", *result)

}
