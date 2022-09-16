package grpc

import (
	"github.com/guil95/split-entrypoints/config/storages/mongo"
	repo "github.com/guil95/split-entrypoints/internal/users/infra/repository/mongo"
	"github.com/guil95/split-entrypoints/internal/users/usecases"
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
		_ = lis.Close()
		zap.S().Errorf("Error to listen server %v", err)
		<-quit
	}

	s := server.NewServer(usecases.NewUseCase(repo.NewCommandMongoRepository(mongo.Connect())))
	grpcServer := grpc.NewServer()
	pb.RegisterUsersServiceServer(grpcServer, s)

	log.Println("Server running on port 50001")
	if err := grpcServer.Serve(lis); err != nil {
		_ = lis.Close()
		zap.S().Errorf("Error in running server %v", err)
		<-quit
	}
}
