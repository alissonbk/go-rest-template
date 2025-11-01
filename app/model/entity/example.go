package entity

// An example entity
type Example struct {
	BaseEntity
	Name string `db:"name" json:"name"`
}
