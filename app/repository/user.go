package repository

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/exception"
	"com.github.alissonbk/go-rest-template/app/model/dto"
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository thus #AutoMigrate will be executed at compile time because of the dependency injection
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) FindAllUser() []*entity.User {
	var users []*entity.User

	rows, err := u.db.Queryx("select * from \"user\"")
	if err != nil {
		logrus.Error("Failed to get all users. Error: ", err)
		exception.PanicException(constant.DBQueryFailed, "")
	}

	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			exception.PanicException(constant.DBQueryFailed, "")
		}
		users = append(users, &entity.User{
			Name:  results["name"].(string),
			Email: results["email"].(string),
		})
	}
	return users
}

func (u UserRepository) Save(user dto.UserDTO) {
	result, err := u.db.Exec(
		`INSERT INTO "user" (name, email, password) VALUES ($1, $2, $3)`,
		user.Name, user.Email, user.Password,
	)
	if err != nil {
		logrus.Error("Failed to save user. Error: ", err.Error())
		exception.PanicException(constant.DBQueryFailed, err.Error())
	}

	if rowsAffected, err := result.RowsAffected(); err != nil || rowsAffected <= 0 {
		logrus.Error("no rows affected")
		exception.PanicException(constant.DBQueryFailed, err.Error())
	}
}

func (u UserRepository) Update(user entity.User) error {
	result := u.db.MustExec("update user set name = $1", user.Name)

	ra, err := result.RowsAffected()
	if err != nil || ra <= 0 {
		return err
	}

	return nil
}

func (u UserRepository) FindUserById(id int) entity.User {
	var user entity.User
	err := u.db.Get(&user, "select * from user where id = $1", id)

	if err != nil {
		exception.PanicException(constant.DBQueryFailed, err.Error())
	}
	return user
}
