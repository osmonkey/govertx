package main

import (
	w "./woo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile("/home/olaf/Development/govertx/justtry/clicert.pem", "local")
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost:49300", opts...)
	if err != nil {

	}
	defer conn.Close()
	client := w.NewWooServiceClient(conn)
	r, err := client.Call(context.Background(), &w.WooRequest{})
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(r.Message)
}
