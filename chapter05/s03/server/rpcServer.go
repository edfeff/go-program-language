package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}
type Quotient struct {
	Quo, Rem int
}

type Arith int

//接口1 乘法
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

//接口2 除法
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

//任意类型即可
type AddSub int

func (a *AddSub) Add(args *Args, result *int) error {
	*result = args.A + args.B
	return nil
}
func (a *AddSub) Sub(args *Args, result *int) error {
	*result = args.A - args.B
	return nil
}

//Start 注册服务对象并开启该 RPC 服务
func Start(addr string) {
	arith := new(Arith)
	addSub := new(AddSub)
	// 注册rpc服务
	rpc.Register(arith)
	rpc.Register(addSub)
	//rpc.Register()
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", addr)
	if e != nil {
		log.Fatal(e)
	}
	go http.Serve(l, nil)
}
