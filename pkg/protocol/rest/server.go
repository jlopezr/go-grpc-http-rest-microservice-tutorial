package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	v1 "github.com/jlopezr/go-grpc-http-rest-microservice-tutorial/pkg/api/v1"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	//swagger := http.FileServer(http.Dir("./3rdparty/swagger-ui"))
	fmt.Println("request", r.URL.Path)
	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	p = path.Join("third-party/swagger-ui/", p)
	fmt.Println("request map ", p)
	http.ServeFile(w, r, p)
}

// RunServer runs HTTP/REST gateway
func RunServer(ctx context.Context, grpcPort, httpPort string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	gw := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := v1.RegisterToDoServiceHandlerFromEndpoint(ctx, gw, "localhost:"+grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/swagger/", serveSwagger)
	mux.Handle("/v1/", gw)

	srv := &http.Server{
		Addr:    ":" + httpPort,
		Handler: mux,
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
		}

		_, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		_ = srv.Shutdown(ctx)
	}()

	log.Println("starting HTTP/REST gateway...")
	return srv.ListenAndServe()
}
