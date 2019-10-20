package main

import (
	"github.com/wpp/go-program-language/chapter05/s03/server"
	"time"
)

//RPC编程
/*
RPC（Remote Procedure Call，远程过程调用）是一种通过网络从远程计算机程序上请求服
务，而不需要了解底层网络细节的应用程序通信协议。RPC协议构建于TCP或UDP，或者是 HTTP
之上，允许开发者直接调用另一台计算机上的程序，而开发者无需额外地为这个调用过程编写网
络通信相关代码，使得开发包括网络分布式程序在内的应用程序更加容易。

RPC 采用客户端—服务器（Client/Server）的工作模式。

Go语言中的RPC支持与处理

在Go中，标准库提供的net/rpc 包实现了 RPC 协议需要的相关细节，开发者可以很方便地
使用该包编写 RPC 的服务端和客户端程序，这使得用 Go 语言开发的多个进程之间的通信变得非
常简单。

net/rpc包允许 RPC 客户端程序通过网络或是其他 I/O 连接调用一个远端对象的公开方法
（必须是大写字母开头、可外部调用的）。在 RPC 服务端，可将一个对象注册为可访问的服务，
之后该对象的公开方法就能够以远程的方式提供访问。一个 RPC 服务端可以注册多个不同类型
的对象，但不允许注册同一类型的多个对象。

一个对象中只有满足如下这些条件的方法，才能被 RPC 服务端设置为可供远程访问：
 必须是在对象外部可公开调用的方法（首字母大写）；
 必须有两个参数，且参数的类型都必须是包外部可以访问的类型或者是Go内建支持的类型；
 第二个参数必须是一个指针；
 方法必须返回一个error类型的值。

以上4个条件，可以简单地用如下一行代码表示：
func (t *T) MethodName(argType T1, replyType *T2) error

类型T、T1 和 T2 默认会使用 Go 内置的 encoding/gob 包进行编码解码

第一个参数表示由 RPC 客户端传入的参数，
第二个参数表示要返回给RPC客户端的结果，
该方法最后返回一个 error 类型的值

RPC 服务端可以通过调用 rpc.ServeConn 处理单个连接请求。

在 RPC 客户端，Go 的 net/rpc 包提供了便利的 rpc.Dial() 和 rpc.DialHTTP() 方法来与指定的 RPC 服务端建立连接
调用 RPC 客户端的 Call() 方法则进行同步处理，
当调用 RPC 客户端的 Go() 方法时，则可以进行异步处理，

如果没有明确指定 RPC 传输过程中使用何种编码解码器，默认将使用 Go 标准库提供的encoding/gob 包进行数据传输。

*/
func main() {
	server.Start(":12345")
	time.Sleep(time.Hour * 1)
}
