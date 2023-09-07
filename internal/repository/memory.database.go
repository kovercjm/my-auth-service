package repository

import "my-auth-service/internal/repository/model"

type database struct {
	tokens map[string]struct{}

	users     map[string]*model.User
	roles     map[string]*model.Role
	userRoles map[string][]*model.UserRole
	trash     []interface{}
}
