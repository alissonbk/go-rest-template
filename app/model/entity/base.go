package entity

import (
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt time.Time `gorm:"->:false;column:deleted_at" json:"-"`
}
