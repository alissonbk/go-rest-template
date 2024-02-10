package repository

import (
	"com.github.alissonbk/go-rest-template/app/model/entity"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Error("Failed to migrate user. Error: ", err)
	}
	return &UserRepository{db: db}
}

func (u UserRepository) FindAllUser() ([]entity.User, error) {
	var users []entity.User

	var err = u.db.Find(&users).Error
	if err != nil {
		log.Error("Failed to get all users. Error: ", err)
		return nil, err
	}

	return users, nil
}

func (u UserRepository) Save(user *entity.User) (entity.User, error) {
	err := u.db.Save(user).Error
	if err != nil {
		log.Error("Failed to save user. Error: ", err)
		return entity.User{}, err
	}
	return *user, nil
}

func (u UserRepository) Update(user *entity.User) error {
	err := u.db.Model(user).Updates(*user).Error
	if err != nil {
		log.Error("Failed to update user. Error: ", err)
		return err
	}
	return nil
}

func (u UserRepository) FindUserById(id int) (entity.User, error) {
	user := entity.User{
		Id: id,
	}
	err := u.db.First(&user, "id = ?", id).Error
	if err != nil {
		log.Error("Failed to find user by id. Error: ", err)
		return entity.User{}, err
	}
	return user, nil
}

func (u UserRepository) DeleteUserById(id int) error {
	err := u.db.Delete(&entity.User{}, id).Error
	if err != nil {
		log.Error("Failed to delete user. Error: ", err)
		return err
	}
	return nil
}
