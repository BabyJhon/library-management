package service

import (
	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type User interface {
	CreateUser(user entity.User) (int, error) //возвращает id созданного в базе пользователя
	GetUser(id int) (entity.User, error)
	DeleteUser(id int) error
}

type Book interface {
	CreateBook(book entity.Book) (int, error)
	DeleteBook(id int) error
	GetBook(id int) (entity.Book, error)
	GetAllBooks() ([]entity.Book, error)
	GiveBookToUser(userId, bookId int) (int, error)
	ReturnBookFromUser(userId, bookId int) error
	GetBooksByUser(userId int) ([]entity.Book, error)
}

type Service struct {
	Book
	User
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		User: NewUserService(repos),
		Book: NewBookService(repos),
	}
}
