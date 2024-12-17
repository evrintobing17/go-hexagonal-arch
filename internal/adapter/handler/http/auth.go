package http

import (
	"net/http"

	"github.com/evrintobing17/go-hexagonal-arch/internal/core/port"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	svc port.AuthService
}

func NewAuthHandler(svc port.AuthService) *AuthHandler {
	return &AuthHandler{
		svc: svc,
	}
}

type AuthRequest struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" valdate:"required"`
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req AuthRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {

	}

	token, err := h.svc.Login(ctx, req.UserName, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, token)

}
