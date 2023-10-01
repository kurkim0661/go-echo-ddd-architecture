package repository

import (
	"myapp/api/domain"
	"myapp/api/domain/dto"
)

type CardRepository interface {
	Save(*dto.CardDto) (*domain.Card, error)
	DeleteById(int, int) error
}
