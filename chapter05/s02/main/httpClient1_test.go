package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"
)

// http基本方法
//func (c *Client) Get(url string) (r *Response, err error)
//func (c *Client) Post(url string, bodyType string, body io.Reader) (r *Response, err error)
//func (c *Client) PostForm(url string, data url.Values) (r *Response, err error)
//func (c *Client) Head(url string) (r *Response, err error)
//func (c *Client) Do(req *Request) (resp *Response, err error)
func Test_GET(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	defer resp.Body.Close()
	//
	io.Copy(os.Stdout, resp.Body)
}

func Test_POST(t *testing.T) {
	// 请求的目标 URL
	// 将要 POST 数据的资源类型（MIMEType）
	// 数据的比特流（[]byte形式）
	resp, err := http.Post("http://127.0.0.1:80", "application/json", nil)
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
	}
	t.Log(resp.Body)
}
func Test_POSTFORM(t *testing.T) {
	resp, err := http.PostForm("http://example.com/posts",
		url.Values{
			"title":   {"article title"},
			"content": {"article body"}})
	if err != nil {
		// 处理错误
		return
	}
	//
	io.Copy(os.Stdout, resp.Body)
}
func Test_HEAD(t *testing.T) {
	resp, err := http.Head("http://127.0.0.1")
	if err != nil {
		// 处理错误
		return
	}
	//
	for k, v := range resp.Header {
		fmt.Println(k, v)
	}
}

//我们希望设定一些自定义的 Http Header 字段
// 设定自定义的"User-Agent"，而不是默认的 "Go http package"
// 传递 Cookie
//此时可以使用net/http包http.Client对象的Do()方法来实现：
func Test_SpecifyHttp(t *testing.T) {
	request, e := http.NewRequest("GET", "http://example.com", nil)
	if e == nil {
		request.Header.Add("User-Agent", "WPP")
		client := &http.Client{}
		response, e := client.Do(request)

		if e == nil {
			defer response.Body.Close()
			io.Copy(os.Stdout, response.Body)
		}
	}
}

//高级封装
//自定义http.Client
//前面我们使用的http.Get()、http.Post()、http.PostForm()和http.Head()方法其
//实都是在http.DefaultClient的基础上进行调用的，比如http.Get()等价于http.Default-
//Client.Get()，依次类推。

//、、http.Client类型的结构：
/*
type Client struct {
	// Transport用于确定HTTP请求的创建机制。
	// 如果为空，将会使用DefaultTransport
	Transport RoundTripper

	// CheckRedirect定义重定向策略。
	// 如果CheckRedirect不为空，客户端将在跟踪HTTP重定向前调用该函数。
	// 两个参数req和via分别为即将发起的请求和已经发起的所有请求，最早的
	// 已发起请求在最前面。
	// 如果CheckRedirect返回错误，客户端将直接返回错误，不会再发起该请求。
	// 如果CheckRedirect为空，Client将采用一种确认策略，将在10个连续
	// 请求后终止
	CheckRedirect func(req *Request, via []*Request) error

	// 如果Jar为空，Cookie将不会在请求中发送，并会
	// 在响应中被忽略
	Jar CookieJar
}
*/
func Test_SpecifyClientCheckRedirect(t *testing.T) {
	//自定义了重定向策略
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("")
			return nil
		},
	}
	resp, err := client.Get("XXX")
	t.Log(resp, err)

	request, err := http.NewRequest("GET", "XXX", nil)
	t.Log(request, err)

	request.Header.Add("User-Agent", "WPP Custom")
	request.Header.Add("If-None-Match", `W/TheFileEtag`)

	resp, err = client.Do(request)
	t.Log(resp, err)
}

//http.Transport 类型的具体结构：
/**
type Transport struct {
	// Proxy指定用于针对特定请求返回代理的函数。
	// 如果该函数返回一个非空的错误，请求将终止并返回该错误。
	// 如果Proxy为空或者返回一个空的URL指针，将不使用代理
	Proxy func(*Request) (*url.URL, error)

	// Dial指定用于创建TCP连接的dail()函数。
	// 如果Dial为空，将默认使用net.Dial()函数
	Dial func(net, addr string) (c net.Conn, err error)

	// TLSClientConfig指定用于tls.Client的TLS配置。
	// 如果为空则使用默认配置
	TLSClientConfig *tls.Config

	//是否取消长连接，默认值为 false，即启用长连接
	DisableKeepAlives bool

	//是否取消压缩（GZip），默认值为 false，即启用压缩
	DisableCompression bool

	// 如果MaxIdleConnsPerHost为非零值，它用于控制每个host所需要
	// 保持的最大空闲连接数。如果该值为空，则使用DefaultMaxIdleConnsPerHost
	MaxIdleConnsPerHost int

	// ...
}
*/
func Test_SpecifyClient(t *testing.T) {
	var pool *x509.CertPool = nil
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Get("XXX")
	t.Log(resp, err)
}
