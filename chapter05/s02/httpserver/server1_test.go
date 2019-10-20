package httpserver

import (
	"fmt"
	"html"
	"net/http"
	"testing"
)

//HTTP服务端

//1. 处理HTTP请求

//使用 net/http 包提供的 http.ListenAndServe() 方法，可以在指定的地址进行监听，
//开启一个HTTP，

//func ListenAndServe(addr string, handler Handler) error
//第一个参数 addr 即监听地址；
// 第二个参数表示服务端处理程序,通常为空，这意味着服务端调用 http.DefaultServeMux 进行处理，
// 而服务端编写的业务逻辑处理程序 http.Handle() 或 http.HandleFunc() 默认注入 http.DefaultServeMux 中
type FooHandler struct {
	Name string
}

func (f *FooHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	//fmt.Println(resp, req)
	resp.WriteHeader(200)
	resp.Header().Add("Content-Type", "text/html")
	resp.Write([]byte("<h1>Hello ServerHttp</h1>"))
}

func TestServer1(t *testing.T) {
	fooHandler := &FooHandler{Name: "wpp-server"}
	//处理实例
	http.Handle("/foo", fooHandler)
	//处理函数
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	//log.Fatal(http.ListenAndServe(":8080", nil))

	server := &http.Server{
		Addr: "localhost:4000",
	}
	t.Fatal(server.ListenAndServe())
	//如果想更多地控制服务端的行为，可以自定义 http.Server，代码如下：
	//s := &http.Server{
	//Addr: ":8080",
	//Handler: myHandler,
	//ReadTimeout: 10 * time.Second,
	//WriteTimeout: 10 * time.Second,
	//MaxHeaderBytes: 1 << 20,
	//}
}
