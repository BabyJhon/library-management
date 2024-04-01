package repo

import (
	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/jmoiron/sqlx"
)

type Book interface {
	CreateBook(book entity.Book) (int, error)
	DeleteBook(id int) error
	GetBook(id int) (entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	AddBookToUser(userId, bookId int) (int, error)
	DeleteBookFromUser(userId, bookId int) error
	GetBooksByUser(userId int) ([]entity.Book, error)
	isBookInLibrary(bookId int) (bool, error)
}

type User interface {
	CreateUser(user entity.User) (int, error)
	GetUser(id int) (entity.User, error)
	DeleteUser(id int) error
}

type Repository struct {
	Book
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		User: NewUserRepo(db),
		Book: NewBookRepo(db),
	}
}
