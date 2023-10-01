package application

import (
	"myapp/api/domain"
	"myapp/api/domain/repository"
)

type UserServiceImpl struct {
	repository.UserRepository
}

func (UserServiceImpl) NewUserServiceImpl(userRepository repository.UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository}
}

func (userServiceImpl *UserServiceImpl) GetUsers() ([]*domain.User, error) {
	return userServiceImpl.UserRepository.FindAll()
}

func (userServiceImpl *UserServiceImpl) CreateUser(user *domain.User) (*domain.User, error) {
	return userServiceImpl.UserRepository.Save(user)
}

func (userServiceImpl *UserServiceImpl) DeleteUser(id int) error {

	// transaction 처리가 필요한 모든 메서드에, tx 객체를 전달
	// 메서드 아규먼트에 tx 추가 vs middleware(echo ctx 공유)를 쓴다 vs 레파지토리에 짠다.
	err := userServiceImpl.UserRepository.DeleteById(id)
	if err != nil {
		//tx.Rollback()
		return err
	}

	return nil
}

func (userServiceImpl *UserServiceImpl) GetUser(id int) (*domain.User, error) {
	return userServiceImpl.UserRepository.FindById(id)
}

func (userServiceImpl *UserServiceImpl) UpdateUser(user *domain.User) (*domain.User, error) {
	return userServiceImpl.UserRepository.Update(user)
}
