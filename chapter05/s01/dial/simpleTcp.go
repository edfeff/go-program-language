package dial

import (
	"fmt"
	"net"
	"os"
)

//tcp示例
func TestTCP() {
	if len(os.Args) != 2 {
		fmt.Println()
		os.Exit(1)
	}
	service := os.Args[1]
	conn, e := net.Dial("tcp", service)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	_, e = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}

	result, e := readFully(conn)
	if e != nil {
		fmt.Println(e)
		os.Exit(1)
	}
	fmt.Println(string(result))
}
