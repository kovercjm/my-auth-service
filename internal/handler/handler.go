package handler

import (
	kLog "github.com/kovercjm/tool-go/logger"

	"my-auth-service/internal/handler/gen"
	"my-auth-service/internal/middleware/dependency"
)

var _ gen.UserAPI = (*Handler)(nil)
var _ gen.RoleAPI = (*Handler)(nil)
var _ gen.AuthAPI = (*Handler)(nil)

type Handler struct {
	logger     kLog.Logger
	repository dependency.Repository
}

func New(logger kLog.Logger, repository dependency.Repository) *Handler {
	return &Handler{
		logger:     logger,
		repository: repository,
	}
}
