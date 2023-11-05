package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/studious-barnacle/auth-api/handler"
	"github.com/ilmsg/studious-barnacle/auth-api/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("auth-api")
	authGRPCURL := os.Getenv("authGRPCURL")

	conn, err := grpc.Dial(authGRPCURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// /auth/profile	{ first_name, last_name, birthday}, jwt token
	client := proto.NewAuthClient(conn)
	hRegister := handler.NewHandlerRegister(client)
	hLogin := handler.NewHandlerLogin(client)

	app := gin.Default()

	app.POST("/register", hRegister.PostRegister) // /auth/register 	{ email, password }
	app.POST("/login", hLogin.PostLogin)          // /auth/login		{ email, password }

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
