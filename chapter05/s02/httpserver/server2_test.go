package httpserver

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"testing"
)

//HTTP服务端
//2. 处理HTTPS请求

//net/http 包还提供 http.ListenAndServeTLS() 方法，用于处理 HTTPS 连接请求：
// func ListenAndServeTLS(addr string, certFile string, keyFile string, handler Handler) error

//服务器上必须存在包含证书和与之匹配的私钥的相关文件，
// 比如certFile对应SSL证书文件存放路径，keyFile对应证书私钥文件路径。
// 如果证书是由证书颁发机构签署的，certFile参数指定的路径必须是存放在服务器上的经由CA认证过的SSL证书

func TestHTTPS(t *testing.T) {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil))
}
