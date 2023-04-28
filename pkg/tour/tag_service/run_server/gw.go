package run_server

import (
	"context"
	"giao/pkg/tour/tag_service/pkg/swagger"
	"giao/pkg/tour/tag_service/proto"
	"giao/pkg/tour/tag_service/server"
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"
)

func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

func RunServer(port string) error {
	httpMux := runHttpServer()
	grpcS := runGrpcServer()
	gatewayMux := runGrpcGatewayServer(port)

	httpMux.Handle("/", gatewayMux)

	// return http.ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
	return ListenAndServe(":"+port, grpcHandlerFunc(grpcS, httpMux))
}

func runHttpServer() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`pong`))
	})

	prefix := "/swagger-ui/"
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    swagger.Asset,
		AssetDir: swagger.AssetDir,
		Prefix:   "third_party/swagger-ui",
	})
	serveMux.Handle(prefix, http.StripPrefix(prefix, fileServer))
	serveMux.HandleFunc("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "swagger.json") {
			http.NotFound(w, r)
			return
		}

		p := strings.TrimPrefix(r.URL.Path, "/swagger/")
		p = path.Join("proto", p)

		http.ServeFile(w, r, p)
	})

	return serveMux
}

func runGrpcServer() *grpc.Server {
	s := grpc.NewServer()
	proto.RegisterTagServiceServer(s, server.NewTagServe())
	reflection.Register(s)

	return s
}

func runGrpcGatewayServer(port string) *runtime.ServeMux {
	endpoint := "0.0.0.0:" + port
	gwmux := runtime.NewServeMux()
	dopts := []grpc.DialOption{grpc.WithInsecure()}
	_ = proto.RegisterTagServiceHandlerFromEndpoint(context.Background(), gwmux, endpoint, dopts)

	return gwmux
}

type tcpKeepAliveListener struct {
	*net.TCPListener
}

func ListenAndServe(addr string, handler http.Handler) error {
	srv := &http.Server{Addr: addr, Handler: handler}
	addr = srv.Addr
	if addr == "" {
		addr = ":http"
	}
	go func(addr string) {
		sign := make(chan os.Signal, 1)
		signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
		data := <-sign

		log.Printf("received sign :%s", data.String())
		log.Println("shutdown...")

		err := srv.Shutdown(context.Background())
		if err != nil {
			log.Fatalf("shutdown %s err:%s", addr, err)
		}
	}(addr)

	l, err := net.Listen("tcp4", addr)
	if err != nil {
		return err
	}
	log.Printf("listen on %s", addr)
	return srv.Serve(tcpKeepAliveListener{l.(*net.TCPListener)})
}
