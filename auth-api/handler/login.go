package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/studious-barnacle/auth-api/proto"
)

type handlerLogin struct {
	client proto.AuthClient
}

func NewHandlerLogin(client proto.AuthClient) *handlerLogin {
	return &handlerLogin{client}
}

func (h *handlerLogin) PostLogin(ctx *gin.Context) {
	var authLogin proto.LoginRequest
	if err := ctx.ShouldBindJSON(&authLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req := &proto.LoginRequest{Email: authLogin.Email, Password: authLogin.Password}
	if res, err := h.client.Login(context.Background(), req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"token": res.Token,
		})
	}
}
