package main

import (
	"context"
	"database/sql"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math"
	"muzz_challenge/cmd/internal/repo"
	"muzz_challenge/cmd/internal/server"
	explore "muzz_challenge/pkg/proto"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	// TODO - init the relevant databases, repository layer and serve the grpc server
	grpcServer := grpc.NewServer()

	// TODO - update the init func
	exploreDB, err := NewMysqlConnection("admin:pass@tcp(host.docker.internal:33062)/explore?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Failed to connect to MySQL: %v", err)
	}

	r := repo.New(exploreDB)
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

func NewMysqlConnection(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to database: %v", err)
	}

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		err = db.Ping()
		if err == nil {
			break
		}

		log.Printf("Database ping attempt %d/%d failed: %v", i+1, maxRetries, err)

		if i < maxRetries-1 {
			waitTime := time.Duration(math.Pow(2, float64(i))) * time.Second
			log.Printf("Waiting %v before next attempt...", waitTime)
			time.Sleep(waitTime)
		}
	}

	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("Failed to ping database after %d attempts: %v", maxRetries, err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
