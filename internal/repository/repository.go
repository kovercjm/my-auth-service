package repository

import (
	kLog "github.com/kovercjm/tool-go/logger"

	"my-auth-service/internal/middleware/dependency"
	"my-auth-service/internal/repository/model"
)

var _ dependency.Repository = (*Repository)(nil)

type Repository struct {
	logger kLog.Logger
	database
}

func New(logger kLog.Logger) (dependency.Repository, error) {
	db := database{
		users:     map[string]*model.User{},
		roles:     map[string]*model.Role{},
		userRoles: map[string][]*model.UserRole{},
		trash:     []interface{}{},
		tokens:    map[string]struct{}{},
	}

	return &Repository{logger: logger, database: db}, nil
}
