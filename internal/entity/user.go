package entity

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" db:"name"`
	Surname  string `json:"surname" db:"surname"`
	Username string `json:"username" db:"username"`
	Phone    string `json:"phone" db:"phone"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password_hash"`
}
