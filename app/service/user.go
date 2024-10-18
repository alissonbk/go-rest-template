package service

import (
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"com.github.alissonbk/go-rest-template/app/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) GetAll() []*entity.User {
	return s.repository.FindAllUser()
}

func (s *UserService) Save(user dto.UserDTO) {
	s.repository.Save(user)
}

func (s *UserService) GetByID(id int) entity.User {
	user := s.repository.FindUserById(id)
	return user
}

func (s *UserService) Update(user entity.User) {
	s.repository.Update(user)
}
