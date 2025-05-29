package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"muzz_challenge/cmd/internal/repo"
	"muzz_challenge/cmd/internal/server"
	explore "muzz_challenge/pkg/proto"
	"net"
	"os"
	"os/signal"
)

func main() {
	// TODO - init the relevant databases, repository layer and serve the grpc server
	grpcServer := grpc.NewServer()

	// TODO - update the init func
	r := repo.New()
	explore.RegisterExploreServiceServer(
		grpcServer,
		server.New(r),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	serve(ctx, grpcServer, cancel)
}

func serve(
	ctx context.Context,
	grpcServer *grpc.Server,
	cancel context.CancelFunc,
) {
	listenAddress := "127.0.0.1:50001"
	listener, listenerErr := net.Listen("tcp", listenAddress)
	if listenerErr != nil {
		log.Fatalf("failed to listen: %v", listenerErr)
	}

	reflection.Register(grpcServer)

	// Run our server in a goroutine so that it doesn't block listening for shutdown
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			cancel()
		}
	}()

	<-ctx.Done()

	// Gracefully shut down gRPC server
	grpcServer.GracefulStop()

	// Close TCP listener
	_ = listener.Close()
}
