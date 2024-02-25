package dto

import (
	"com.github.alissonbk/go-rest-template/app/model/entity"
)

type UserDTO struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (dto UserDTO) ToEntity() entity.User {
	return entity.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
