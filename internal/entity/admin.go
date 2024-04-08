package entity

type Admin struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required" db:"name"`
	SureName string `json:"sure_name" binding:"required" db:"sure_name"`

	UserName string `json:"user_name" binding:"required" db:"user_name"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}
