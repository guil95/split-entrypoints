package server

import (
	"context"

	pb "github.com/guil95/split-entrypoints/proto/genpb/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
}

func NewServer() pb.UsersServiceServer {
	return &server{}
}

func (s server) GetUsers(ctx context.Context, empty *emptypb.Empty) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{Users: nil}, nil
}
