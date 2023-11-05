package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/ilmsg/studious-barnacle/auth-grpc/database"
	"github.com/ilmsg/studious-barnacle/auth-grpc/proto"
	"github.com/ilmsg/studious-barnacle/auth-grpc/repo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AuthServer struct {
	proto.AuthServer
	repo *repo.RepoAuth
}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	mongoURL := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")
	db := database.NewMongoDatabase(mongoURL, dbName)

	collectionName := os.Getenv("collectionName")
	collection := db.GetMongoCollection(collectionName)

	repo := repo.NewRepoAuth(collection)
	authServer := &AuthServer{repo: repo}
	// proto.RegisterAuthServer(srv, &AuthServer{})

	srv := grpc.NewServer()
	proto.RegisterAuthServer(srv, authServer)
	reflection.Register(srv)

	if err := srv.Serve(lst); err != nil {
		log.Fatal(err)
	}
}

func (a *AuthServer) Register(ctx context.Context, req *proto.RegisterRequest) (*proto.LoginResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()
	fmt.Println("register:")
	fmt.Printf("email: %s\n", email)
	fmt.Printf("password: %s\n", password)

	token, err := a.repo.Register(email, password)
	if err != nil {
		return &proto.LoginResponse{Token: ""}, err
	}
	return &proto.LoginResponse{Token: token}, nil
}

func (a *AuthServer) Login(ctx context.Context, req *proto.LoginRequest) (*proto.LoginResponse, error) {
	email, password := req.GetEmail(), req.GetPassword()

	fmt.Println("login:")
	fmt.Printf("email: %s\n", email)
	fmt.Printf("password: %s\n", password)

	token, err := a.repo.Login(email, password)
	if err != nil {
		return &proto.LoginResponse{Token: ""}, err
	}
	return &proto.LoginResponse{Token: token}, nil
}
