package dial

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
)

func TestDialICMP() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[1]
	fmt.Println("before send to", service)
	conn, err := net.Dial("ip4:icmp", service)
	fmt.Println("connect is send")
	checkError(err)
	var msg [512]byte
	msg[0] = 8  //echo
	msg[1] = 0  //code 0
	msg[2] = 0  //checksum
	msg[3] = 0  //checksum
	msg[4] = 0  //identifier[0]
	msg[5] = 13 //identifier[1]
	msg[6] = 0  //sequence[0]
	msg[7] = 37 //sequence[1]
	len := 8
	check := checkSum(msg[0:len])
	msg[2] = byte(check >> 8)
	msg[3] = byte(check & 255)

	_, err = conn.Write(msg[0:len])
	checkError(err)

	_, err = conn.Read(msg[0:])
	checkError(err)

	fmt.Println("GOT response")
	if msg[5] == 13 {
		fmt.Println("Identifier matches")
	}
	if msg[5] == 17 {
		fmt.Println("Sequence matches")
	}
	os.Exit(0)
}
func checkSum(msg []byte) uint16 {
	sum := 0
	//假设为偶数
	for n := 1; n < len(msg)-1; n += 2 {
		sum += int(msg[n])*256 + int(msg[n+1])
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += (sum >> 16)
	var a uint16 = uint16(^sum)
	return a
}
func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal Error:%s", err.Error())
		os.Exit(1)
	}
}
func readFully(conn net.Conn) ([]byte, error) {
	defer conn.Close()
	result := bytes.NewBuffer(nil)
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
	}
	return result.Bytes(), nil
}
