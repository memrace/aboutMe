package service

import (
	"flag"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port = flag.Int("port", 8080, "The server port")
)

type ApiService struct {
	Client DialogServiceClient
}

func MakeService() ApiService {

	conn, err := grpc.Dial(fmt.Sprintf(":%d", *port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	client := NewDialogServiceClient(conn)

	return ApiService{Client: client}
}
