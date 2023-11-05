package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ilmsg/studious-barnacle/auth-api/proto"
)

type handlerRegister struct {
	client proto.AuthClient
}

func NewHandlerRegister(client proto.AuthClient) *handlerRegister {
	return &handlerRegister{client}
}

func (h *handlerRegister) PostRegister(ctx *gin.Context) {
	var authRegis proto.LoginRequest
	if err := ctx.ShouldBindJSON(&authRegis); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req := &proto.RegisterRequest{Email: authRegis.Email, Password: authRegis.Password}
	if res, err := h.client.Register(context.Background(), req); err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"token": res.Token,
		})
	}
}
