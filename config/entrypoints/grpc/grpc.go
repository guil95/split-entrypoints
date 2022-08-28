package grpc

import (
	"log"
	"net"
	"os"

	"github.com/guil95/split-entrypoints/internal/users/infra/server/grpc/server"
	pb "github.com/guil95/split-entrypoints/proto/genpb/users"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func RunGrpcServer(quit chan os.Signal) {
	lis, err := net.Listen("tcp", "0.0.0.0:50001")
	if err != nil {
		zap.S().Errorf("Error to listen server %v", err)
		<-quit
	}

	s := server.NewServer()
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, s)

	log.Println("Server running on port 50001")
	if err := grpcServer.Serve(lis); err != nil {
		zap.S().Errorf("Error in running server %v", err)
		<-quit
	}
}
