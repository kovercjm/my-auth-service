package model

import "time"

type User struct {
	ID           string
	Name         string
	PasswordHash string
	timestamp
}

type Role struct {
	ID   string
	Name string
	timestamp
}

type UserRole struct {
	UserID string
	RoleID string
	timestamp
}

type timestamp struct {
	CreatedAt time.Time
	DeletedAt *time.Time
}
