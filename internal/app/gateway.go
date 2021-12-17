package app

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"price_analyzer_prototype/api"
	"price_analyzer_prototype/internal/price_analyzer_api"
)

func GRPCService(serviceAddr string, analyzerAPI *price_analyzer_api.ProductAnalyzerAPI) {
	lis, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		log.Fatalln("failed to listen TCP port", err)
	}
	server := grpc.NewServer()

	api.RegisterPriceAnalyzeServer(server, analyzerAPI)

	fmt.Println("starting gRPC server at " + serviceAddr)
	server.Serve(lis)
}

func HTTPProxy(proxyAddr, serviceAddr string) {
	grpcConn, err := grpc.Dial(
		serviceAddr,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalln("failed to connect to grpc", err)
	}
	defer grpcConn.Close()

	grpcGWMux := runtime.NewServeMux()
	err = api.RegisterPriceAnalyzeHandler(
		context.Background(),
		grpcGWMux,
		grpcConn,
	)
	if err != nil {
		log.Fatalln("failed to start HTTP server", err)
	}

	//client := api.NewPriceAnalyzeClient(grpcConn)
	mux := http.NewServeMux()

	mux.Handle("/api/v1/products", grpcGWMux)
	mux.Handle("/api/v1/categories", grpcGWMux)
	mux.Handle("/api/v1/filters", grpcGWMux)
	mux.HandleFunc("/", versionHandler)

	fmt.Println("stating HTTP server at " + proxyAddr)
	log.Fatalln(http.ListenAndServe(proxyAddr, mux))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "version")
	if err != nil {
		return
	}
}
