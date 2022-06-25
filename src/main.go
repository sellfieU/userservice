package main

import (
	"context"
	"github.com/sellfie/usermanagerservice/src/pb"
	"google.golang.org/grpc"
)

func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServiceServer{})
}

type userServiceServer struct{}

func (u userServiceServer) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceServer) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceServer) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}

func (u userServiceServer) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}
