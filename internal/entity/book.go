package entity

type Book struct {
	Id        int    `json:"id" db:"id"`
	Title     string `json:"title" binding:"required" db:"title"`
	Author    string `json:"author" binding:"required" db:"author"`
	InLibrary bool   `json:"in_library" db:"in_library"`
}
