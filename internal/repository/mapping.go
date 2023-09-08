package repository

import (
	"my-auth-service/internal/domain"
	"my-auth-service/internal/repository/model"
	"my-auth-service/internal/service"
)

func userToModel(user *domain.User) (*model.User, error) {
	passwordHash, err := service.Hash(user.Password)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:           user.ID,
		Name:         user.Name,
		PasswordHash: passwordHash,
	}, nil
}

func roleToModel(role *domain.Role) *model.Role {
	return &model.Role{
		ID:   role.ID,
		Name: role.Name,
	}
}

func userRoleToModel(userRole *domain.UserRole) *model.UserRole {
	return &model.UserRole{
		UserID: userRole.User.ID,
		RoleID: userRole.Roles[0].ID,
	}
}

func userRoleToDomain(userID string, userRole []*model.UserRole) *domain.UserRole {
	if len(userRole) == 0 {
		return nil
	}
	// TODO get actual names
	result := &domain.UserRole{User: &domain.User{ID: userID, Name: userID}}
	for i := 0; i < len(userRole); i++ {
		result.Roles = append(result.Roles, &domain.Role{ID: userRole[i].RoleID, Name: userRole[i].RoleID})
	}
	return result
}
