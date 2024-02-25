package entity

import (
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time `gorm:"not null; ->:false; column:created_at; default:current_timestamp" json:"-"`
	UpdatedAt time.Time `gorm:"->:false;column:updated_at" json:"-"`
	DeletedAt time.Time `gorm:"->:false;column:deleted_at" json:"-"`
}
