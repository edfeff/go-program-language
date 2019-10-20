package dial

import (
	"fmt"
	"net"
)

//在Go语言中编写网络程序时，我们将看不到传统的编码形式。以前我们使用Socket编程时，
//会按照如下步骤展开。
//(1) 建立Socket：使用socket()函数。
//(2) 绑定Socket：使用bind()函数。
//(3) 监听：使用listen()函数。或者连接：使用connect()函数。
//(4) 接受连接：使用accept()函数。
//(5) 接收：使用receive()函数。或者发送：使用send()函数。
//Go语言标准库对此过程进行了抽象和封装。无论我们期望使用什么协议建立什么形式的连
//接，都只需要调用net.Dial()即可。

//func Dial(net, addr string) (Conn, error) {}
//net参数是网络协议的名字
//addr参数是IP地址或域名，而端口号以“:”的形式跟随在地址或域名的后面，端口号可选。
// 如果连接成功，返回连接对象，否则返回error。

//tcp
func tcpDial() {
	conn, e := net.Dial("tcp", "x.x.x.x:8080")
	fmt.Println(conn, e)
}

//udp
func udpDial() {
	conn, e := net.Dial("udp", "x.x.x.x:975")
	fmt.Println(conn, e)
}

//icmp
func icmpDial() {
	conn, e := net.Dial("ip4:icmp", "www.baidu.com")
	fmt.Println(conn, e)
}

//使用协议编号进行指定协议
//协议编号参考
//https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
func icmpDial2() {
	conn, e := net.Dial("ip4:1", "10.0.0.3")
	fmt.Println(conn, e)
}

//icmp程序示例
