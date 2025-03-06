package handler

import (
	"Hexagonal-Architecture/internal/core/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) LoginUser(ctx *gin.Context) {
	var user domain.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	response, err := h.svc.LoginUser(user.Email, user.Password)
	if err != nil {
		HandleError(ctx, http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":            response.ID,
		"email":         response.Email,
		"access_token":  response.AccessToken,
		"refresh_token": response.RefreshToken,
		"is_member":     response.Membership,
	})
}
