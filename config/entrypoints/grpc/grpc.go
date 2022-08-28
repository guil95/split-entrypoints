package grpc

import (
	"log"
	"net"

	"github.com/guil95/split-entrypoints/internal/users/infra/server/grpc/server"
	pb "github.com/guil95/split-entrypoints/proto/genpb/users"
	"google.golang.org/grpc"
)

func RunGrpcServer(quit chan error) {
	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		quit <- err
	}

	s := server.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, s)

	log.Println("Server running on port 50001")
	if err := grpcServer.Serve(lis); err != nil {
		quit <- err
	}
}
