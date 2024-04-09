package entity

type User struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required" db:"name"` 
	SureName    string `json:"sure_name" binding:"required" db:"sure_name"`
	PhoneNumber string `json:"phone_number" binding:"required" db:"phone_number"`
}
