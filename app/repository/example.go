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
