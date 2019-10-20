package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//6.5 Go 语言的哈希函数
//Go提供了MD5、SHA-1等几种哈希函数，
func hash1() {
	s := "hello world"
	md5Inst := md5.New()
	md5Inst.Write([]byte(s))
	result := md5Inst.Sum([]byte(""))
	fmt.Println(result)        //[94 182 59 187 224 30 238 208 147 203 34 187 143 90 205 195]
	fmt.Printf("%x\n", result) //5eb63bbbe01eeed093cb22bb8f5acdc3

	sha1Inst := sha1.New()
	sha1Inst.Write([]byte(s))
	result = sha1Inst.Sum([]byte(""))
	fmt.Println(result)        //[42 174 108 53 201 79 207 180 21 219 233 95 64 139 156 233 30 232 70 237]
	fmt.Printf("%x\n", result) //2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
}

//文件哈希
func hash2() {
	dir, _ := filepath.Abs(`.`)
	fName := dir + "\\Test1.txt"
	file, e := os.Open(fName)
	if e == nil {
		md5h := md5.New()
		io.Copy(md5h, file)
		fmt.Printf("%x\n", md5h.Sum([]byte(""))) //0b4e7a0e5fe84ad35fb5f95b9ceeac79
		sha1h := sha1.New()
		io.Copy(sha1h, file)
		fmt.Printf("%x\n", sha1h.Sum([]byte(""))) //da39a3ee5e6b4b0d3255bfef95601890afd80709
	} else {
		fmt.Println(e)
	}
}
