package dto

import (
	"myapp/api/domain"
)

type UserDto struct {
	Name string `json:"name"`
}

func (UserDto *UserDto) ToEntity() *domain.User {
	return &domain.User{Name: UserDto.Name}
}

type CardDto struct {
	Name   string `json:"name"`
	Limit  int    `json:"limit"`
	UserId uint   `json:"userId"`
}

func (cardDto *CardDto) ToEntity() *domain.Card {
	return &domain.Card{Name: cardDto.Name, Limit: cardDto.Limit}
}
