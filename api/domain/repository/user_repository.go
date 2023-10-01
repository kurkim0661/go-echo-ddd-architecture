package repository

import (
	"myapp/api/domain"
)

type UserRepository interface {
	FindAll() ([]*domain.User, error)
	FindById(id int) (*domain.User, error)
	DeleteById(id int) error
	Save(user *domain.User) (*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
}
