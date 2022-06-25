package userservice

import (
	"context"
	"github.com/go-logr/logr"
	"github.com/sellfie/usermanagerservice/src/pb"
)

// UserService is a struct that have user service gRPC methods
type UserService struct {
	Log logr.Logger
}

// CreateUser is a method to create new user
func (u *UserService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}

// GetUser is a method to get a user's info
func (u *UserService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	//TODO implement me
	panic("implement me")
}

// DeleteUser is a method to delete user
func (u *UserService) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}

// LoginUser is a method to set a user status as logged in
func (u *UserService) LoginUser(ctx context.Context, request *pb.LoginUserRequest) (*pb.Error, error) {
	//TODO implement me
	panic("implement me")
}
