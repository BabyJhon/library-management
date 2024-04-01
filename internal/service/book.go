package service

import (
	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/internal/repo"
)

type BookService struct {
	repo repo.Book
}

func NewBookService(repo repo.Book) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) CreateBook(book entity.Book) (int, error) {
	book.InLibrary = true //добавление книги происходит с учетом, что оа уже в библиотеке
	return b.repo.CreateBook(book)
}

func (b *BookService) DeleteBook(id int) error {
	return b.repo.DeleteBook(id)
}

func (b *BookService) GetBook(id int) (entity.Book, error) {
	return b.repo.GetBook(id)
}

func (b *BookService) GetAllBooks() ([]entity.Book, error) {
	return b.repo.GetAllBooks()
}

func (b *BookService) GiveBookToUser(userId, bookId int) (int, error) {
	return b.repo.AddBookToUser(userId, bookId)
}

func (b *BookService) ReturnBookFromUser(userId, bookId int) error {
	return b.repo.DeleteBookFromUser(userId, bookId)
}

func (b *BookService) GetBooksByUser(userId int) ([]entity.Book, error) {
	return b.repo.GetBooksByUser(userId)
}
