package application

import (
	"myapp/api/domain"
	"myapp/api/domain/dto"
	"myapp/api/domain/repository"
)

type CardServiceImpl struct {
	repository.CardRepository
}

func (CardServiceImpl) NewCardServiceImpl(repository repository.CardRepository) *CardServiceImpl {
	return &CardServiceImpl{repository}
}

func (cardServiceImpl *CardServiceImpl) CreateCard(cardDto *dto.CardDto) (*domain.Card, error) {
	return cardServiceImpl.CardRepository.Save(cardDto)
}

func (cardServiceImpl *CardServiceImpl) DeleteCard(cardId int, userId int) error {
	return cardServiceImpl.CardRepository.DeleteById(cardId, userId)
}
