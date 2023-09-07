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

func (h Handler) RolesPost(ctx *gin.Context) {
	request := &gen.RoleInfo{}
	if err := ctx.BindJSON(request); err != nil {
		return
	}

	if err := h.repository.CreateRole(&domain.Role{
		ID:   request.Name,
		Name: request.Name,
	}); err != nil {
		response := &kAPI.Error{}
		switch {
		case errors.Is(err, dependency.AlreadyExistsError):
			response.HTTPStatus = http.StatusConflict
			response.Message = "Role already exists"
		default:
			h.logger.Error("Create role failed", "error", err)
		}
		_ = ctx.Error(response)
		return
	}

	ctx.AbortWithStatusJSON(http.StatusOK, &gen.UsersPost200Response{Id: request.Name})
}

func (h Handler) RolesIdDelete(ctx *gin.Context) {
	roleID := ctx.Param("id")
	if err := h.repository.DeleteRole(roleID); err != nil {
		response := &kAPI.Error{}
		switch {
		case errors.Is(err, dependency.NotFoundError):
			response.HTTPStatus = http.StatusNotFound
			response.Message = "Role not exists"
		default:
			h.logger.Error("Delete role failed", "error", err)
		}
		_ = ctx.Error(response)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
