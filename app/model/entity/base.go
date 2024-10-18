package entity

import (
	"time"
)

type BaseEntity struct {
	ID        uint      `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	DeletedAt time.Time `db:"deleted_at"`
}
