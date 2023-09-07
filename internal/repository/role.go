package repository

import (
	"time"

	"my-auth-service/internal/domain"
	"my-auth-service/internal/middleware/dependency"
)

func (r Repository) CreateRole(role *domain.Role) error {
	if role == nil {
		return dependency.InvalidArgumentError
	}
	if _, ok := r.database.roles[role.ID]; ok {
		return dependency.AlreadyExistsError
	}

	record := roleToModel(role)
	record.CreatedAt = time.Now()
	r.database.roles[record.ID] = record

	return nil
}

func (r Repository) DeleteRole(id string) error {
	record, ok := r.database.roles[id]
	if !ok {
		return dependency.NotFoundError
	}

	delete(r.database.roles, id)

	deleteAt := time.Now()
	record.DeletedAt = &deleteAt
	r.database.trash = append(r.database.trash, record)

	return nil
}
