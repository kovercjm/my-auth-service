package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	kAPI "github.com/kovercjm/tool-go/server/api"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/handler/gen"
	"my-auth-service/internal/middleware/dependency"
)

func (h Handler) UsersPost(ctx *gin.Context) {
	request := &gen.UserInfo{}
	if err := ctx.BindJSON(request); err != nil {
		return
	}

	if err := h.repository.CreateUser(&domain.User{
		ID:       request.Name,
		Name:     request.Name,
		Password: request.Password,
	}); err != nil {
		response := &kAPI.Error{}
		switch {
		case errors.Is(err, dependency.AlreadyExistsError):
			response.HTTPStatus = http.StatusConflict
			response.Message = "User already exists"
		default:
			h.logger.Error("Create user failed", "error", err)
		}
		_ = ctx.Error(response)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, &gen.UsersPost200Response{Id: request.Name})
}

func (h Handler) UsersIdDelete(ctx *gin.Context) {
	userID := ctx.Param("id")
	if err := h.repository.DeleteUser(userID); err != nil {
		response := &kAPI.Error{}
		switch {
		case errors.Is(err, dependency.NotFoundError):
			response.HTTPStatus = http.StatusNotFound
			response.Message = "User not exists"
		default:
			h.logger.Error("Delete user failed", "error", err)
		}
		_ = ctx.Error(response)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
