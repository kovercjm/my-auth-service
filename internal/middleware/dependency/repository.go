package dependency

import (
	"github.com/pkg/errors"

	"my-auth-service/internal/domain"
)

type Repository interface {
	CreateUser(user *domain.User) error
	CheckUserPassword(user *domain.User) (bool, error)
	DeleteUser(id string) error

	CreateRole(role *domain.Role) error
	DeleteRole(id string) error

	GrantUserRole(userRole *domain.UserRole) error
	GetUserRoles(user *domain.User) (*domain.UserRole, error)

	CreateToken(token string) error
	CheckToken(token string) error
	DeleteToken(token string) error
}

var (
	AlreadyExistsError   = errors.New("already exists")
	InvalidArgumentError = errors.New("invalid argument")
	NotFoundError        = errors.New("not found")
)
