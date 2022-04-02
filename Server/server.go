package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	pb "server/protos/user"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type userAuthServer struct {
	pb.UnimplementedUserAuthServer
}

func (s *userAuthServer) Register(ctx context.Context, in *pb.RegReq) (*pb.ApiRes, error) {
	// if username is not in database, and everything is alright, register, and store in database
	// else, return error message to client that username is already in database or something is wrong
	log.Printf("Registered: %v", in.GetUsername())
	return &pb.ApiRes{ResCode: 200, Message: "success"}, nil
}

func (s *userAuthServer) Login(ctx context.Context, in *pb.LoginReq) (*pb.ApiRes, error) {
	// if username is in database, and password is correct, login, and return success message to client
	// else, return error message to client that username is not in database or password is wrong
	log.Printf("Logged in: %v", in.GetUsername())
	return &pb.ApiRes{ResCode: 200, Message: "success"}, nil
}

func (s *userAuthServer) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.ApiRes, error) {
	return &pb.ApiRes{ResCode: 200, Message: "successfully logged out"}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserAuthServer(s, &userAuthServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
