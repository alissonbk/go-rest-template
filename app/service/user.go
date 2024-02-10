package service

import (
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"com.github.alissonbk/go-rest-template/app/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAll() ([]entity.User, error) {
	users, err := s.repository.FindAllUser()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) Save(user entity.User) (entity.User, error) {
	savedUser, err := s.repository.Save(&user)
	if err != nil {
		return entity.User{}, err
	}
	return savedUser, nil
}

func (s *UserService) GetByID(id int) (entity.User, error) {
	user, err := s.repository.FindUserById(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *UserService) Update(user entity.User) error {
	return s.repository.Update(&user)
}

func (s *UserService) Delete(id int) error {
	return s.repository.DeleteUserById(id)
}
