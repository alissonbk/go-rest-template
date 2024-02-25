package repository

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/exception"
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository thus #AutoMigrate will be executed at compile time because of the dependency injection
func NewUserRepository(db *gorm.DB) *UserRepository {
	err := db.AutoMigrate(&entity.User{})
	if err != nil {
		panic("Failed to migrate user: " + err.Error())
	}
	return &UserRepository{db: db}
}

func (u UserRepository) FindAllUser() []entity.User {
	var users []entity.User

	var err = u.db.Find(&users).Error
	fmt.Println(users)
	if err != nil {
		log.Error("Failed to get all users. Error: ", err)
		exception.PanicException(constant.DBQueryFailed, "")
	}

	return users
}

func (u UserRepository) Save(user *entity.User) entity.User {
	tx := u.db.Save(user)
	if tx.Error != nil {
		log.Error("Failed to save user. Error: ", tx.Error)
		if strings.Contains(tx.Error.Error(), "duplicate key") {
			exception.PanicException(constant.DBDuplicatedKey, "Email already exists")
		}
		exception.PanicException(constant.DBQueryFailed, "")
	}
	return *user
}

func (u UserRepository) Update(user entity.User) {
	log.Info(user)

	tx := u.db.Model(&user).Updates(user)

	if tx.RowsAffected < 1 {
		log.Warning("0 Rows affected.")
		exception.PanicException(constant.DBNoRowsAffected, "")
	}
	if tx.Error != nil {
		log.Error("Failed to update user. Error: ", tx.Error)
		exception.PanicException(constant.DBQueryFailed, "xd")
	}
}

func (u UserRepository) FindUserById(id int) entity.User {
	user := entity.User{
		Id: id,
	}
	tx := u.db.First(&user)
	if tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "record not found") {
			exception.PanicException(constant.DataNotFound, "Not found")
		}
		log.Error("Failed to find user by id. Error: ", tx.Error)
		exception.PanicException(constant.DBQueryFailed, "")
	}
	return user
}

func (u UserRepository) DeleteUserById(id int) {
	tx := u.db.Delete(&entity.User{}, id)

	if tx.RowsAffected < 1 {
		log.Warn("Delete got 0 rows affected")
		exception.PanicException(constant.DBNoRowsAffected, "No rows affected")
	}
	if tx.Error != nil {
		log.Error("Failed to delete user. Error: ", tx.Error)
		exception.PanicException(constant.DBQueryFailed, "")
	}
}
