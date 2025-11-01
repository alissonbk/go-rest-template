package dto

import (
	"com.github.alissonbk/go-rest-template/app/model/entity"
)

type ExampleDTO struct {
	Name string `json:"name"`
}

func (dto ExampleDTO) ToEntity() entity.Example {
	return entity.Example{
		Name: dto.Name,
	}
}
