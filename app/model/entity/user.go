package entity

import "com.github.alissonbk/go-rest-template/app/constant"

type User struct {
	Name     string        `gorm:"column:name" json:"name"`
	Email    string        `gorm:"column:email;index:idx_email,unique" json:"email"`
	Password string        `gorm:"column:password;->:false" json:"-"`
	Status   int           `gorm:"column:status" json:"status"`
	Role     constant.Role `gorm:"column:role" json:"role"`
	BaseEntity
}
