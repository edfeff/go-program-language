package main

import (
	"net/http"
	"testing"
)

//灵活的 http.RoundTripper 接口
/*
type RoundTripper interface {
	// RoundTrip执行一个单一的HTTP事务，返回相应的响应信息。
	// RoundTrip函数的实现不应试图去理解响应的内容。如果RoundTrip得到一个响应，
	// 无论该响应的HTTP状态码如何，都应将返回的err设置为nil。非空的err
	// 只意味着没有成功获取到响应。
	// 类似地，RoundTrip也不应试图处理更高级别的协议，比如重定向、认证和
	// Cookie等。
	//
	// RoundTrip不应修改请求内容, 除非了是为了理解Body内容。每一个请求
	// 的URL和Header域都应被正确初始化
	RoundTrip(*Request) (*Response, error)
}
*/
type OurTrans struct {
	Transport http.RoundTripper
}

func (o *OurTrans) transport() http.RoundTripper {
	if o.Transport != nil {
		return o.Transport
	}
	return http.DefaultTransport
}
func (o *OurTrans) RoundTrip(req *http.Request) (*http.Response, error) {
	// 处理一些事情 ...
	// 发起HTTP请求
	// 添加一些域到req.Header中
	return o.transport().RoundTrip(req)
}
func (o *OurTrans) Client() *http.Client {
	return &http.Client{Transport: o}
}

//因为实现了http.RoundTripper 接口的代码通常需要在多个 goroutine中并发执行，因此我
//们必须确保实现代码的线程安全性。
func Test_SpecifyRoundTripper(t *testing.T) {
	o := &OurTrans{
		Transport: nil,
	}

	c := o.Client()
	resp, err := c.Get("http://example.com")
	// ...
	t.Log(resp, err)
}
