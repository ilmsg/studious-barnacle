package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	pb "github.com/ilmsg/studious-barnacle/auth-proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("auth-api")

	conn, err := grpc.Dial("localhost:7001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	// /auth/login		{ email, password }
	// /auth/profile	{ first_name, last_name, birthday}, jwt token
	client := pb.NewAuthClient(conn)

	app := gin.Default()

	// /auth/register 	{ email, password }
	app.POST("/register", func(ctx *gin.Context) {
		var authRegis pb.LoginRequest
		if err := ctx.ShouldBindJSON(&authRegis); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		req := &pb.RegisterRequest{Email: authRegis.Email, Password: authRegis.Password}
		if res, err := client.Register(context.Background(), req); err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"token": res.Token,
			})
		}
	})

	if err := app.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
