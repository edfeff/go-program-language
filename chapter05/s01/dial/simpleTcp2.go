package dial

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

//此外，net包中还包含了一系列的工具函数，合理地使用这些函数可以更好地保障程序的质量。

//验证IP地址有效性的代码如下：
//func net.ParseIP()

//创建子网掩码的代码如下：
//func IPv4Mask(a, b, c, d byte) IPMask

//获取默认子网掩码的代码如下：
//func (ip IP) DefaultMask() IPMask

//根据域名查找IP的代码如下：
//func ResolveIPAddr(net, addr string) (*IPAddr, error)
//func LookupHost(name string) (cname string, addrs []string, err error)；

func TestTCP2() {
	if len(os.Args) != 2 {
		fmt.Println()
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service) //用于解析地址和端口号

	if err != nil {
		fmt.Println(err)
		return
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr) //用于建立链接
	//这两个函数在Dial()中都得到了封装
	if err != nil {
		fmt.Println(err)
		return
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(result))
}
