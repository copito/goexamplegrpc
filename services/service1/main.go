package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"

	// Import path from go_package option (just look into go.mod in idl_gen)
	pb "github.com/copito/example_grpc/idl_gen/go/service1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// server is used to implement service1.HelloServiceServer
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements service1.HelloServiceServer
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
	slog.Info("started building server")
	// Start the HTTP server for the gRPC-Gateway in a goroutine
	go func() {
		// Create a gRPC-Gateway mux
		mux := runtime.NewServeMux()

		// Register the HelloService handler from the generated code
		// Route registration (for rest counterpart)
		err := pb.RegisterGreeterHandlerServer(context.Background(), mux, &server{})
		if err != nil {
			slog.Error("failed to register gRPC-Gateway handler", slog.String("err", err.Error()))
		}

		// Create a HTTP server for the gRPC-Gateway
		gwServer := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}

		slog.Info("Starting gRPC-Gateway on port 8080")
		if err := gwServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Info("failed to start gRPC-Gateway", slog.String("err", err.Error()))
		}
	}()

	// Start the gRPC server
	go func() {
		// Create a listener on TCP port 50051
		lis, err := net.Listen("tcp", ":50051")
		if err != nil {
			slog.Error("failed to listen", slog.String("err", err.Error()))
			panic("unable to start connection for grpc on port 50051")
		}

		// Create a gRPC server object
		grpcServer := grpc.NewServer()

		// Register the HelloService service with the server
		// Route registration (like rest)
		pb.RegisterGreeterServer(grpcServer, &server{})

		// Register reflection service on gRPC server
		reflection.Register(grpcServer)

		slog.Info("Starting gRPC server on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to serve", slog.String("err", err.Error()))
		}
	}()

	select {}
}
