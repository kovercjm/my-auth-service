package repository

import (
	"time"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/middleware/dependency"
)

func (r Repository) GrantUserRole(userRole *domain.UserRole) error {
	if userRole == nil || userRole.User == nil || userRole.Roles == nil || len(userRole.Roles) != 1 {
		return dependency.InvalidArgumentError
	}
	userID, roleID := userRole.User.ID, userRole.Roles[0].ID
	if _, ok := r.database.users[userID]; !ok {
		return dependency.NotFoundError
	}
	if _, ok := r.database.roles[roleID]; !ok {
		return dependency.NotFoundError
	}

	userRoles := r.database.userRoles[userID]
	for _, exist := range userRoles {
		if exist.RoleID == roleID {
			return nil
		}
	}

	record := userRoleToModel(userRole)
	record.CreatedAt = time.Now()
	r.database.userRoles[userID] = append(userRoles, record)

	return nil
}

func (r Repository) GetUserRoles(user *domain.User) (*domain.UserRole, error) {
	if user == nil {
		return nil, dependency.InvalidArgumentError
	}

	userRoles := r.database.userRoles[user.ID]
	if len(userRoles) == 0 {
		return &domain.UserRole{User: &domain.User{ID: user.ID}}, nil
	}

	return userRoleToDomain(user.ID, userRoles), nil
}

func (r Repository) CreateToken(token string) error {
	if token == "" {
		return dependency.InvalidArgumentError
	}
	r.tokens[token] = struct{}{}
	return nil
}

func (r Repository) CheckToken(token string) error {
	if token == "" {
		return dependency.InvalidArgumentError
	}
	if _, ok := r.tokens[token]; !ok {
		return dependency.NotFoundError
	}
	return nil
}

func (r Repository) DeleteToken(token string) error {
	if token == "" {
		return dependency.InvalidArgumentError
	}
	delete(r.tokens, token)
	return nil
}
