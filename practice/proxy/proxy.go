package proxy

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type Pxy struct{}

func (p *Pxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("received request %s %s %s\n", req.Method, req.Host, req.RemoteAddr)
	transPort := http.DefaultTransport
	// 代理接收到客户端的请求，复制了原来的请求对象，并根据数据配置新请求的各种参数（添加上 X-Forward-For 头部等）
	outReq := new(http.Request)
	*outReq = *req
	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := outReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + "," + clientIP
		}
		outReq.Header.Set("X-Forwarded-For", clientIP)
	}
	// 2 把新请求发送到服务器端，并接收到服务器端返回的响应
	res, err := transPort.RoundTrip(outReq)
	if err != nil {
		fmt.Println("transPort err", err)
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	// 3 代理服务器对响应做一些处理，然后返回给客户端
	for k, i := range res.Header {
		for _, v := range i {
			rw.Header().Add(k, v)
		}
	}
	rw.WriteHeader(res.StatusCode)
	_, err = io.Copy(rw, res.Body)
	if err != nil {
		fmt.Println("write err", err)
		return
	}

	err = req.Body.Close()
	if err != nil {
		fmt.Println("Close err", err)
		return
	}
}

// TestProxy 需要在客户端配置代理服务器
func TestProxy() {
	host := "0.0.0.0"
	port := ":9090"
	fmt.Println("serve on ", host, port)
	http.Handle("/", &Pxy{})
	http.ListenAndServe(host+port, nil)
}

// ReverseProxy 反向代理
func ReverseProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.URL.Path = target.Path
		fmt.Println("target", target)
	}
	return &httputil.ReverseProxy{Director: director}
}

func TestReverseProxy() {
	urls := []*url.URL{
		{Scheme: "http", Host: "127.0.0.1:8000"},
		{Scheme: "http", Host: "127.0.0.1:8080"}, // vue 页面js无法加载
	}
	proxy := ReverseProxy(urls)

	http.ListenAndServe(":9090", proxy)
}
