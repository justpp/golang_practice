package run_server

import (
	"giao/src/tour/tag_service/proto"
	"giao/src/tour/tag_service/server"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

func CMux(port string) {
	l, err := RunTcpServer(port)
	if err != nil {
		log.Fatalf("run tcp err:%s", err)
	}

	m := cmux.New(l)
	grpcl := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldPrefixSendSettings("content-type", "application/grpc"))
	httpl := m.Match(cmux.HTTP1Fast())

	grpcS := RunRpcServer()
	httpS := RunHttpServer(port)
	go grpcS.Serve(grpcl)
	go httpS.Serve(httpl)

	err = m.Serve()
	if err != nil {
		log.Fatalf("run cmux err:%s", err)
	}

}

func RunTcpServer(port string) (net.Listener, error) {
	return net.Listen("tcp", ":"+port)
}

func RunHttpServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})
	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}

func RunRpcServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServe())
	reflection.Register(s)
	return s
}
