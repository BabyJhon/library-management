package repo

import (
	"context"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Book interface {
	CreateBook(ctx context.Context, book entity.Book) (int, error)
	DeleteBook(ctx context.Context, id int) error
	GetBook(ctx context.Context, id int) (entity.Book, error)
	GetAllBooks(ctx context.Context) ([]entity.Book, error)
	AddBookToUser(ctx context.Context, userId, bookId int) (int, error)
	DeleteBookFromUser(ctx context.Context, userId, bookId int) error
	GetBooksByUser(ctx context.Context, userId int) ([]entity.Book, error)

	isBookInLibrary(ctx context.Context, bookId int) (bool, error)
}

type User interface {
	CreateUser(ctx context.Context, user entity.User) (int, error)
	GetUser(ctx context.Context, id int) (entity.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Repository struct {
	Book
	User
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{
		User: NewUserRepo(db),
		Book: NewBookRepo(db),
	}
}
