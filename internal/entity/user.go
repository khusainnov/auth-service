package entity

type User struct {
	Username   string `json:"username" db:"username"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Patronymic string `json:"patronymic,omitempty" db:"patronymic"`
	Email      string `json:"email" db:"email"`
	Password   string `json:"password" db:"password_hash"`
}
