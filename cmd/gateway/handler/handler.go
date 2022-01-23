package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func RegisterHandler(c *gin.Context) {
	//转发的url，端口
	target := "127.0.0.1:8080"
	u := &url.URL{}
	//转发的协议，如果是https，写https.
	u.Scheme = "http"
	u.Host = target
	proxy := httputil.NewSingleHostReverseProxy(u)
	//重写出错回调
	proxy.ErrorHandler = func(rw http.ResponseWriter, req *http.Request, err error) {
		log.Printf("http: proxy error: %v", err)
		ret := fmt.Sprintf("http proxy error %v", err)
		//写到body里
		rw.Write([]byte(ret))
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
