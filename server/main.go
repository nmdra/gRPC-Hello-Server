package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/nmdra/gRPC-Hello-Server/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
)

type greetServer struct {
	pb.UnimplementedGreetServiceServer
}

func (s *greetServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := fmt.Sprintf("Hello, %s!", req.GetName())
	return &pb.HelloResponse{Message: msg}, nil
}

func (s *greetServer) WhoAmI(ctx context.Context, _ *emptypb.Empty) (*pb.WhoAmIResponse, error) {
	remoteAddr := "unknown"
	if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
		remoteAddr = p.Addr.String()
	}

	userAgent := ""
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if ua, found := md["user-agent"]; found && len(ua) > 0 {
			userAgent = ua[0]
		}
	}

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return &pb.WhoAmIResponse{
		RemoteAddr: remoteAddr,
		UserAgent:  userAgent,
		Hostname:   hostname,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()

	// Register service
	pb.RegisterGreetServiceServer(server, &greetServer{})

	// Enable reflection
	reflection.Register(server)

	log.Println("gRPC server listening on :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
