package entity

import "com.github.alissonbk/go-rest-template/app/constant"

type User struct {
	Id       int           `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string        `gorm:"column:name" json:"name"`
	Email    string        `gorm:"column:email;index:idx_email,unique" json:"email"`
	Password string        `gorm:"column:password;" json:"password"`
	Status   int           `gorm:"column:status" json:"status"`
	Role     constant.Role `gorm:"column:role" json:"role"`
	BaseEntity
}
