package application

import (
	"myapp/api/domain"
)

type UserService interface {
	GetUsers() ([]*domain.User, error)
	CreateUser(*domain.User) (*domain.User, error)
	DeleteUser(int) error
	GetUser(int) (*domain.User, error)
	UpdateUser(*domain.User) (*domain.User, error)
}
