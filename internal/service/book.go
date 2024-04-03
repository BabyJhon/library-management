package service

import (
	"context"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type BookService struct {
	repo repo.Book
}

func NewBookService(repo repo.Book) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) CreateBook(ctx context.Context, book entity.Book) (int, error) {
	book.InLibrary = true //добавление книги происходит с учетом, что она уже в библиотеке
	return b.repo.CreateBook(ctx, book)
}

func (b *BookService) DeleteBook(ctx context.Context, id int) error {
	return b.repo.DeleteBook(ctx, id)
}

func (b *BookService) GetBook(ctx context.Context, id int) (entity.Book, error) {
	return b.repo.GetBook(ctx, id)
}

func (b *BookService) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	return b.repo.GetAllBooks(ctx)
}

func (b *BookService) GiveBookToUser(ctx context.Context, userId, bookId int) (int, error) {
	return b.repo.AddBookToUser(ctx, userId, bookId)
}

func (b *BookService) ReturnBookFromUser(ctx context.Context, userId, bookId int) error {
	return b.repo.DeleteBookFromUser(ctx, userId, bookId)
}

func (b *BookService) GetBooksByUser(ctx context.Context, userId int) ([]entity.Book, error) {
	return b.repo.GetBooksByUser(ctx, userId)
}
