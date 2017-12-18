package main

import (
	w "./woo"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"net"
)

type wServer struct {
}

func (ws *wServer) Call(context.Context, *w.WooRequest) (*w.WooResponse, error) {
	r := w.WooResponse{}
	r.Message = "Hallo"
	return &r, nil
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 49300))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("/home/olaf/Development/govertx/justtry/certificate.pem", "/home/olaf/Development/govertx/justtry/key.pem")
	opts := []grpc.ServerOption{grpc.Creds(creds)}
	grpcServer := grpc.NewServer(opts...)
	ws := wServer{}
	w.RegisterWooServiceServer(grpcServer, &ws)
	log.Println("Server started")
	grpcServer.Serve(lis)
}
