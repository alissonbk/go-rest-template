package entity

// The User domain it's only for example purpose...
type User struct {
	Id       int    `gorm:"column:id; primary_key; not null" json:"id"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email;index:idx_email,unique" json:"email"`
	Password string `gorm:"column:password;" json:"password"`
	IsActive bool   `gorm:"default:true; column:is_active" json:"isActive"`
	BaseEntity
}
