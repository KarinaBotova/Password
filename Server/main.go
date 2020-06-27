package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/KarinaBotova/Password/Server/proto"
)

const (
	port = "9090"
)

type Server struct {

}

func (s *Server) Generate(ctx context.Context, req *proto.PasswordGeneratorReq) (*proto.PasswordGeneratorRes, error) {
	if req.GetLength() < 0 {

		return nil, fmt.Errorf("введено отрицательное число!")
	}
	pass := make([]byte, req.GetLength())

	for i := int32(0); i < req.Length; i++ {
		pass[i] = byte('0' + rand.Int31n(10))
	}

	log.Printf("Пароль: %s\n", pass)

	return &proto.PasswordGeneratorRes{Password: string(pass)}, nil
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen gRPC on port %v: %v", port, err)
	}

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{}))

	proto.RegisterPasswordGeneratorServer(server, &Server{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
}
