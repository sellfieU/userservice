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

package server

import (
	"fmt"
	"github.com/sellfie/usermanagerservice/src/pb"
	"github.com/sellfie/usermanagerservice/src/userservice"
	"google.golang.org/grpc"
	"net"
	"os"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	log = logf.Log.WithName("user-service-server")
)

// Server is an interface of server
type Server interface {
	Start(string)
}

type userServiceServer struct {
	grpcServer  *grpc.Server
	userService *userservice.UserService
}

// New returns new user service gRPC server
func New() Server {
	server := new(userServiceServer)
	server.grpcServer = grpc.NewServer()
	server.userService = &userservice.UserService{Log: log}

	return server
}

// Start starts gRPC server on input port
func (s *userServiceServer) Start(port string) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(err, "cannot launch gRPC server")
		os.Exit(1)
	}

	pb.RegisterUserServiceServer(s.grpcServer, s.userService)
	if err := s.grpcServer.Serve(l); err != nil {
		log.Error(err, "cannot launch gRPC server")
		os.Exit(1)
	}
}
