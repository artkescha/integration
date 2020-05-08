package client

import (
	"context"
	"github.com/artkescha/integration/grpc_api"
	"google.golang.org/grpc"
	"log"
	"time"
)

func Run(ctx context.Context, address string, a uint64, accum *int) {
	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %s", err)
	}
	log.Printf("connect to grpc server: %s", address)

	client := grpc_api.NewAccumulatorClient(connection)

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	result, err := client.Reset(ctx, &grpc_api.Request{})
	if err != nil {
		log.Fatalf("reset failed: %s", err)
	}

	result, err = client.Accumulate(ctx, &grpc_api.Request{
		Addendum: a,
	})
	if err != nil {
		log.Fatalf("accumulated failed: %s", err)
	}

	log.Printf("got result: %d", result.Result)
	*accum = int(result.Result)
}
