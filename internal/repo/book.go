package repo

import (
	"errors"
	"fmt"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/pkg/postgres"
	"github.com/jmoiron/sqlx"
)

type BookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (u *BookRepo) CreateBook(book entity.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, author, in_library) values ($1, $2, $3) RETURNING id", postgres.BooksTable)
	row := u.db.QueryRow(query, book.Title, book.Author, book.InLibrary)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (b *BookRepo) DeleteBook(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgres.BooksTable)
	_, err := b.db.Exec(query, id)
	return err
}

func (b *BookRepo) GetBook(id int) (entity.Book, error) {
	var book entity.Book
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.BooksTable)
	err := b.db.Get(&book, query, id)
	return book, err
}

func (b *BookRepo) GetAllBooks() ([]entity.Book, error) {
	var books []entity.Book
	query := fmt.Sprintf("SELECT * FROM %s", postgres.BooksTable)
	err := b.db.Select(&books, query)
	return books, err
}

func (b *BookRepo) isBookInLibrary(bookId int) (bool, error) {
	var inLib bool
	query := fmt.Sprintf("SELECT in_library FROM %s WHERE id = $1", postgres.BooksTable)
	err := b.db.Get(&inLib, query, bookId)
	return inLib, err
}

func (b *BookRepo) AddBookToUser(userId, bookId int) (int, error) { //вернет id вставки
	tx, err := b.db.Begin()
	if err != nil {
		return 0, err
	}

	inLib, err := b.isBookInLibrary(bookId)
	if err != nil {
		return 0, err
	}
	if !inLib { //книга не в библиотеке
		return 0, errors.New("the book is not in the library")
	}

	setBookNotInLibraryQuery := fmt.Sprintf("UPDATE %s SET in_library = false WHERE id = $1", postgres.BooksTable)
	_, err = tx.Exec(setBookNotInLibraryQuery, bookId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) values ($1, $2) RETURNING id", postgres.UsersBooksTable)

	row := tx.QueryRow(query, userId, bookId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (b *BookRepo) DeleteBookFromUser(userId, bookId int) error {
	tx, err := b.db.Begin()
	if err != nil {
		return err
	}

	inLib, err := b.isBookInLibrary(bookId)
	if err != nil {
		return err
	}
	if inLib {
		return errors.New("the book is in the library")
	}

	setBookInLibraryQuery := fmt.Sprintf("UPDATE %s SET in_library = true WHERE id = $1", postgres.BooksTable)
	_, err = tx.Exec(setBookInLibraryQuery, bookId)
	if err != nil {
		tx.Rollback()
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND book_id = $2", postgres.UsersBooksTable)
	res, err := tx.Exec(query, userId, bookId)

	if err != nil {
		tx.Rollback()
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}
	if count == 0 {
		tx.Rollback()
		return errors.New("the user does not have this book")
	}
	//var deletedCount int //для случая, когда пытаемся удалить книгу у пользователя А, хотя она взята у пользователя Б(книга не в библиотеке, пользователи существуют)
	//QueryRow в таком случае вернет 0 т.к. пользователи книги не совпадудт, а не ошибку, поэтому нужна доп проверка
	// err = row.Scan(&deletedCount);
	// TODO
	// надо реализовать удаление и проверку, что что-то удалилось
	// if err != nil {
	// 	tx.Rollback()
	// 	fmt.Println("privet epta")
	// 	return err
	// }
	// if row == 0 {
	// 	tx.Rollback()
	// 	return errors.New("no rows deleted")
	// }

	return tx.Commit()
}

func (b *BookRepo) GetBooksByUser(userId int) ([]entity.Book, error) {
	var books_id []int
	idQuery := fmt.Sprintf("SELECT book_id FROM %s WHERE user_id = $1", postgres.UsersBooksTable)
	err := b.db.Select(&books_id, idQuery, userId)
	if err != nil {
		return nil, err
	}

	var books []entity.Book
	var book entity.Book
	for _, bookId := range books_id {
		bookQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.BooksTable)
		err := b.db.Get(&book, bookQuery, bookId)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
