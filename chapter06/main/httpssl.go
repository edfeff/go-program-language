package main

import (
	"fmt"
	"net/http"
)

//加密通信

//一般的HTTPS是基于SSL（Secure Sockets Layer）协议。
// SSL是网景公司开发的位于TCP与 HTTP之间的透明安全协议，
// 通过SSL，可以把HTTP包数据以非对称加密的形式往返于浏览器和站点之间，从而避免被第三方非法获取。

//目前，伴随着电子商务的兴起，HTTPS获得了广泛的应用。
// 由IETF（Internet Engineering Task Force）实现的TLS（Transport Layer Security）是建立于SSL v3.0之上的兼容协议，
// 它们主要的区别在于所支持的加密算法。

//6.6.1 加密通信流程
/**
(1) 在浏览器中输入HTTPS协议的网址
(2) 服务器向浏览器返回证书，浏览器检查该证书的合法性，
(3) 验证合法性
(4) 浏览器使用证书中的公钥加密一个随机对称密钥，并将加密后的密钥和使用密钥加密后的请求URL一起发送到服务器。
(5) 服务器用私钥解密随机对称密钥，并用获取的密钥解密加密的请求URL。
(6) 服务器把用户请求的网页用密钥加密，并返回给用户。
(7) 用户浏览器用密钥解密服务器发来的网页数据，并将其显示出来。
上述过程都是依赖于SSL/TLS层实现的。


*/
const SERVER_PORT = 8080
const SERVER_DOMAIN = "localhost"
const RESPONSE_TEMPLATE = "hello"

func rootHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Header().Set("Content-Length", fmt.Sprint(len(RESPONSE_TEMPLATE)))
	w.Write([]byte(RESPONSE_TEMPLATE))
}
func HttpsStart() {
	http.HandleFunc(fmt.Sprintf("%s:%d/", SERVER_DOMAIN, SERVER_PORT), rootHandler)
	http.ListenAndServeTLS(fmt.Sprintf(":%d", SERVER_PORT), "rui.crt", "rui.key", nil)
}
