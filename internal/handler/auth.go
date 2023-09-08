package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	kAPI "github.com/kovercjm/tool-go/server/api"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/handler/gen"
	"my-auth-service/internal/middleware/dependency"
	"my-auth-service/internal/service"
)

func (h Handler) UsersUserIDRolesRoleIDPost(ctx *gin.Context) {
	userID, roleID := ctx.Param("userID"), ctx.Param("roleID")
	if err := h.repository.GrantUserRole(&domain.UserRole{
		User:  &domain.User{ID: userID},
		Roles: []*domain.Role{{ID: roleID}},
	}); err != nil {
		response := &kAPI.Error{}
		switch {
		case errors.Is(err, dependency.NotFoundError):
			response.Message = "Invalid Request Parameter"
		default:
			h.logger.Error("Grant user role failed", "error", err)
		}
		_ = ctx.Error(response)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (h Handler) UsersMeRolesGet(ctx *gin.Context) {
	userID := ctx.GetString("userID")
	userRoles, err := h.repository.GetUserRoles(&domain.User{ID: userID})
	if err != nil {
		h.logger.Error("Get user roles failed", "error", err)
		_ = ctx.Error(&kAPI.Error{})
		return
	}
	fmt.Println(userRoles)
	var roles []gen.RoleInfo
	for _, role := range userRoles.Roles {
		roles = append(roles, gen.RoleInfo{
			Id:   role.ID,
			Name: role.Name,
		})
	}
	ctx.AbortWithStatusJSON(http.StatusOK, &gen.UsersMeRolesGet200Response{Roles: roles})
}

func (h Handler) UsersMeRolesIdGet(ctx *gin.Context) {
	userID, roleID := ctx.GetString("userID"), ctx.Param("id")
	userRoles, err := h.repository.GetUserRoles(&domain.User{ID: userID})
	if err != nil {
		h.logger.Error("Get user roles failed", "error", err)
		_ = ctx.Error(&kAPI.Error{})
		return
	}
	for _, role := range userRoles.Roles {
		if role.ID == roleID {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
	_ = ctx.Error(&kAPI.Error{HTTPStatus: http.StatusForbidden})
}

func (h Handler) AuthLoginPost(ctx *gin.Context) {
	request := &gen.UserInfo{}
	if err := ctx.BindJSON(request); err != nil {
		return
	}

	user := &domain.User{
		ID:       request.Name,
		Name:     request.Name,
		Password: request.Password,
	}
	valid, err := h.repository.CheckUserPassword(user)
	if err != nil || !valid {
		h.logger.Error("Auth user login failed", "error", err)
		_ = ctx.Error(&kAPI.Error{HTTPStatus: http.StatusForbidden})
		return
	}

	token, err := service.SignToken(h.repository, &service.Claims{UserID: user.ID})
	if err != nil {
		h.logger.Error("Sign token failed", "error", err)
		_ = ctx.Error(&kAPI.Error{})
		return
	}
	ctx.AbortWithStatusJSON(http.StatusOK, &gen.AuthLoginPost200Response{Token: token})
}

func (h Handler) AuthLogoutPost(ctx *gin.Context) {
	token := ctx.GetString("token")

	if err := h.repository.DeleteToken(token); err != nil {
		h.logger.Error("Delete token failed", "error", err)
		_ = ctx.Error(&kAPI.Error{})
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
