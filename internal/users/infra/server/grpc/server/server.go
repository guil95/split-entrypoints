package server

import (
	"context"
	"github.com/guil95/split-entrypoints/internal/users"
	"github.com/guil95/split-entrypoints/internal/users/usecases"

	pb "github.com/guil95/split-entrypoints/proto/genpb/users"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	uc usecases.UseCase
}

func NewServer(uc usecases.UseCase) pb.UsersServiceServer {
	return &server{uc}
}

func (s server) GetUsers(ctx context.Context, empty *emptypb.Empty) (*pb.GetUsersResponse, error) {
	return &pb.GetUsersResponse{Users: nil}, nil
}

func (s server) SaveUser(ctx context.Context, user *pb.User) (*pb.SaveUsersResponse, error) {
	err := s.uc.Save(ctx, &users.User{Name: user.Name, Age: user.Age})
	if err != nil {
		return nil, err
	}
	return &pb.SaveUsersResponse{User: &pb.User{Name: user.Name, Age: user.Age}}, nil
}
