package repository

import (
	"com.github.alissonbk/go-rest-template/app/constant"
	"com.github.alissonbk/go-rest-template/app/exception"
	"com.github.alissonbk/go-rest-template/app/model/entity"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type ExampleRepository struct {
	db *sqlx.DB
}

// AutoMigrate will be executed at compile time because of the dependency injection
func NewExampleRepository(db *sqlx.DB) *ExampleRepository {
	return &ExampleRepository{db: db}
}

func (u ExampleRepository) FindAllExample() []*entity.Example {
	var examples []*entity.Example

	rows, err := u.db.Queryx("select * from \"example\"")
	if err != nil {
		logrus.Error("Failed to get all examples. Error: ", err)
		exception.PanicException(constant.DBQueryFailed, "")
	}

	for rows.Next() {
		results := make(map[string]interface{})
		err = rows.MapScan(results)
		if err != nil {
			exception.PanicException(constant.DBQueryFailed, "")
		}
		examples = append(examples, &entity.Example{
			Name: results["name"].(string),
		})
	}
	return examples
}

// package repository

// import (
// 	"com.github.alissonbk/go-rest-template/app/constant"
// 	"com.github.alissonbk/go-rest-template/app/exception"
// 	"com.github.alissonbk/go-rest-template/app/model/entity"
// 	log "github.com/sirupsen/logrus"
// 	"gorm.io/gorm"
// )

// type ExampleRepository struct {
// 	db *gorm.DB
// }

// // AutoMigrate will be executed at compile time because of the dependency injection
// func NewExampleRepository(db *gorm.DB) *ExampleRepository {
// 	err := db.AutoMigrate(&entity.Example{})
// 	if err != nil {
// 		panic("Failed to migrate example: " + err.Error())
// 	}
// 	return &ExampleRepository{db: db}
// }

// func (u ExampleRepository) FindAllUser() []entity.Example {
// 	var examples []entity.Example

// 	var err = u.db.Find(&examples).Error
// 	if err != nil {
// 		log.Error("Failed to get all users. Error: ", err)
// 		exception.PanicException(constant.DBQueryFailed, "")
// 	}

// 	return examples
// }
