// Copyright 2022 Sellfie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
