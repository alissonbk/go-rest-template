package entity

import (
	"time"
)

type BaseEntity struct {
	ID        int       `gorm:"column:id; primary_key; not null" json:"id"`
	CreatedAt time.Time `gorm:"->:false;column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt time.Time `gorm:"->:false;column:deleted_at" json:"-"`
}
