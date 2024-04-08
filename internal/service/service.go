package service

import (
	"context"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error) //возвращает id созданного в базе пользователя
	GetUser(ctx context.Context, id int) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Book interface {
	CreateBook(ctx context.Context, book entity.Book) (int, error)
	DeleteBook(ctx context.Context, id int) error
	GetBook(ctx context.Context, id int) (entity.Book, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	GiveBookToUser(ctx context.Context, userId, bookId int) (int, error)
	ReturnBookFromUser(ctx context.Context, userId, bookId int) error
	GetBooksByUser(ctx context.Context, userId int) ([]entity.Book, error)
}

type Authorization interface {
	CreateAdmin(ctx context.Context, admin entity.Admin) (int, error)
	GenerateToken(ctx context.Context, userName, password string) (string, error)
	Parsetoken(token string) (int, error)
}

type Service struct {
	Book
	User
	Authorization
}

func NewService(repos *repo.Repository) *Service {
	return &Service{
		User:          NewUserService(repos),
		Book:          NewBookService(repos),
		Authorization: NewAuthService(repos),
	}
}
