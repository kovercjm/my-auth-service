package repository

import (
	"fmt"
	"time"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/middleware/dependency"
	"my-auth-service/internal/service"
)

func (r Repository) CreateUser(user *domain.User) error {
	if user == nil {
		return dependency.InvalidArgumentError
	}
	if _, ok := r.database.users[user.ID]; ok {
		return dependency.AlreadyExistsError
	}

	record, err := userToModel(user)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}
	record.CreatedAt = time.Now()
	r.database.users[record.ID] = record

	return nil
}

func (r Repository) CheckUserPassword(user *domain.User) (bool, error) {
	record, ok := r.database.users[user.ID]
	if !ok {
		return false, dependency.NotFoundError
	}

	passwordHash, err := service.Hash(user.Password)
	if err != nil {
		return false, err
	}
	return record.PasswordHash == passwordHash, nil
}

func (r Repository) DeleteUser(id string) error {
	deleteAt := time.Now()

	userRecord, ok := r.database.users[id]
	if !ok {
		return dependency.NotFoundError
	}
	delete(r.database.users, id)

	userRecord.DeletedAt = &deleteAt
	r.database.trash = append(r.database.trash, userRecord)

	userRoleRecord, ok := r.database.userRoles[id]
	if ok {
		delete(r.database.userRoles, id)
		for _, userRole := range userRoleRecord {
			userRole.DeletedAt = &deleteAt
			r.database.trash = append(r.database.trash, userRole)
		}
	}

	return nil
}
