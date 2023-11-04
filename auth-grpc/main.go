package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/ilmsg/studious-barnacle/auth-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct {
	pb.AuthServer
}

func main() {
	lst, err := net.Listen("tcp", ":7001")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	pb.RegisterAuthServer(srv, &AuthServer{})
	reflection.Register(srv)

	if err := srv.Serve(lst); err != nil {
		log.Fatal(err)
	}
}

func (a *AuthServer) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.LoginResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()
	fmt.Println("register:")
	fmt.Printf("email: %s\n", email)
	fmt.Printf("password: %s\n", password)

	token := ""
	return &pb.LoginResponse{Token: token}, nil
}

func (a *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()

	fmt.Println("login:")
	fmt.Printf("email: %s\n", email)
	fmt.Printf("password: %s\n", password)

	token := ""
	return &pb.LoginResponse{Token: token}, nil
}
