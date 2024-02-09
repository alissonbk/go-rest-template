package constant

type Role int

const (
	Admin Role = iota
	User
)

func (r Role) GetRole() string {
	return [2]string{"ADMIN", "USER"}[r]
}
