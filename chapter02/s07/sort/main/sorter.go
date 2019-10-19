package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/wpp/go-program-language/chapter02/s07/sort/main/bubblesort"
	"github.com/wpp/go-program-language/chapter02/s07/sort/main/qsort"
	"io"
	"os"
	"strconv"
	"time"
)

var infile *string = flag.String("i", "infile", " file to sorting")
var outfile *string = flag.String("o", "outfile", " outfile")
var algorithm *string = flag.String("a", "algorithm", " algorithm")

//读取文件
func readValues(in string) (values []int, err error) {
	file, err := os.Open(in)
	if err != nil {
		fmt.Printf("open file:%s error:%s", in, err)
		return
	}
	defer file.Close()
	br := bufio.NewReader(file)
	values = make([]int, 0)
	//逐行读取
	for {
		line, isPrefix, e := br.ReadLine()

		//处理文件读取结束和读取错误
		if e != nil {
			//文件结束标识
			if e != io.EOF {
				//保存错误，退出一起送出
				err = e
			}
			//文件结束标识后退出
			break
		}
		if isPrefix {
			fmt.Println("too long line!")
			return
		}
		str := string(line) //ReadLine读取出的是[]byte 需要转成字符串，在转数字
		//转成int
		value, e2 := strconv.Atoi(str)
		if e2 != nil {
			err = e2
			return
		}
		values = append(values, value)
	}
	return
}

//写入文件
func writeValues(values []int, out string) error {
	file, err := os.Create(out)
	if err != nil {
		fmt.Printf("create file %s failed", out)
		return err
	}
	//确保文件关闭
	defer file.Close()
	for _, v := range values {
		//转成字符串
		s := strconv.Itoa(v)
		file.WriteString(s + "\n")
	}
	return nil
}
func main() {
	//解析命令行参数
	flag.Parse()
	if infile != nil {
		fmt.Printf("infile: %s\noutfile:%s\nalgorithm:%s\n", *infile, *outfile, *algorithm)
	}
	values, err := readValues(*infile)
	if err != nil {
		fmt.Println(err)
	} else {
		//fmt.Println(values)
		t1 := time.Now()
		switch *algorithm {
		case "b":
			bubblesort.BubbleSort(values)
		case "q":
			qsort.QuickSort(values)
		default:
			qsort.QuickSort(values)
		}
		t2 := time.Now()
		fmt.Printf("sort finish ,time:%d", t2.Sub(t1))
		writeValues(values, *outfile)
	}
}
