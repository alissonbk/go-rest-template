package entity

// The User domain it's only for example purpose...
type User struct {
	BaseEntity
	Name     string `db:"name" json:"name"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
	IsActive bool   `db:"is_active" json:"isActive"`
}
