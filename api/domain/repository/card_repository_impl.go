package repository

import (
	"myapp/api/domain"
	"myapp/api/domain/dto"

	"gorm.io/gorm"
)

type CardRepositoryImpl struct {
	db *gorm.DB
}

func (CardRepositoryImpl) NewCardRepositoryImpl(db *gorm.DB) *CardRepositoryImpl {
	return &CardRepositoryImpl{db}
}

func (cardRepositoryImpl *CardRepositoryImpl) Save(cardDto *dto.CardDto) (*domain.Card, error) {
	card := cardDto.ToEntity()
	err := cardRepositoryImpl.db.Save(card).Error
	if err != nil {
		return nil, err
	}
	return card, nil
}

func (cardRepositoryImpl *CardRepositoryImpl) GetCards() (*[]domain.Card, error) {
	cards := &[]domain.Card{}
	err := cardRepositoryImpl.db.Table("cards").Find(&cards).Error
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func (cardRepositoryImpl *CardRepositoryImpl) DeleteById(cardId int, userId int) error {
	err := cardRepositoryImpl.db.Where("id = ?", cardId).Where("user_id = ?", userId).Delete(&domain.Card{}).Error
	if err != nil {
		return err
	}
	return nil
}
