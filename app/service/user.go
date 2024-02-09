package service

import "com.github.alissonbk/go-rest-template/app/repository"

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(repository *repository.UserRepository) *UserService {
	return &UserService{repository: repository}
}
