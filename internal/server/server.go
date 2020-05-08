package server

import (
	"context"
	"github.com/artkescha/integration/grpc_api"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync/atomic"
)

type Accumulate struct {
	memory uint64
}

func (a *Accumulate) Accumulate(ctx context.Context, req *grpc_api.Request) (*grpc_api.Response, error) {
	atomic.AddUint64(&a.memory, req.Addendum)
	return &grpc_api.Response{Result: a.memory}, nil
}

func (a *Accumulate) Reset(ctx context.Context, req *grpc_api.Request) (*grpc_api.Response, error) {
	a.memory = 0
	return &grpc_api.Response{Result: a.memory}, nil
}

func Run(ctx context.Context, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen %s", err)
	}

	server := grpc.NewServer()
	grpc_api.RegisterAccumulatorServer(server, new(Accumulate))
	log.Printf("start grpc server at %s", address)
	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("failed to serve %s", err)
	}
}
