package service

import (
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"com.github.alissonbk/go-rest-template/app/repository"
)

type ExampleService struct {
	repository *repository.ExampleRepository
}

func NewExampleService(repository *repository.ExampleRepository) *ExampleService {
	return &ExampleService{repository: repository}
}

func (s *ExampleService) GetAll() []*entity.Example {
	return s.repository.FindAllExample()
}
