package custom_net

import (
	"net"
	"net/http"
)

func ListenAndServe(Addr string, handler http.Handler) error {
	srv := &http.Server{Addr: Addr, Handler: handler}
	addr := srv.Addr
	if addr == "" {
		addr = ":http"
	}
	// 指定tcp6 默认 tcp会同时开启ipv4和ipv6
	ln, err := net.Listen("tcp6", addr)
	if err != nil {
		return err
	}
	return srv.Serve(ln)
}
