package application

import (
	"myapp/api/domain"
	"myapp/api/domain/dto"
)

type CardService interface {
	CreateCard(*dto.CardDto) (*domain.Card, error)
	DeleteCard(int, int) error
}
