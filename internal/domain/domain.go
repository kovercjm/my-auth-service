package domain

type User struct {
	ID       string
	Name     string
	Password string
}

type Role struct {
	ID   string
	Name string
}

type UserRole struct {
	User  *User
	Roles []*Role
}
