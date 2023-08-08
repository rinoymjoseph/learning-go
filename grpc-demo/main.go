package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"rinoymjoseph.github.com/grpc-demo/calc"
)

type calcServer struct {
	calc.UnimplementedCalcServer
}

func (s calcServer) Add(ctx context.Context, req *calc.AddRequest) (*calc.AddResponse, error) {
	return &calc.AddResponse{
		Sum: int64(req.Num1 + req.Num2),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatalf("cannot create listener: %s", err)
	}
	serverRegistrar := grpc.NewServer()
	service := &calcServer{}

	calc.RegisterCalcServer(serverRegistrar, service)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %s", err)
	}
}
