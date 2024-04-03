package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/BabyJhon/library-managment/internal/entity"
	"github.com/BabyJhon/library-managment/pkg/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (u *BookRepo) CreateBook(ctx context.Context, book entity.Book) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title, author, in_library) values ($1, $2, $3) RETURNING id", postgres.BooksTable)

	row := u.db.QueryRow(ctx, query, book.Title, book.Author, book.InLibrary)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (b *BookRepo) DeleteBook(ctx context.Context, id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", postgres.BooksTable)
	row, err := b.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected := row.RowsAffected()
	if rowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}

func (b *BookRepo) GetBook(ctx context.Context, id int) (entity.Book, error) {
	var book entity.Book
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.BooksTable)

	row := b.db.QueryRow(ctx, query, id)
	if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.InLibrary); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.Book{}, ErrBookNotFound
		}
		return entity.Book{}, err
	}

	return book, nil
}

func (b *BookRepo) GetAllBooks(ctx context.Context) ([]entity.Book, error) {
	var books []entity.Book
	query := fmt.Sprintf("SELECT * FROM %s", postgres.BooksTable)

	rows, err := b.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book entity.Book
		if err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.InLibrary); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepo) isBookInLibrary(ctx context.Context, bookId int) (bool, error) {
	var inLib bool
	query := fmt.Sprintf("SELECT in_library FROM %s WHERE id = $1", postgres.BooksTable)

	row := b.db.QueryRow(ctx, query, bookId)
	if err := row.Scan(&inLib); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, ErrBookNotFound
		}
		return false, err
	}
	return inLib, nil
}

func (b *BookRepo) AddBookToUser(ctx context.Context, userId, bookId int) (int, error) { //вернет id вставки
	tx, err := b.db.Begin(ctx)
	if err != nil {
		return 0, err
	}

	inLib, err := b.isBookInLibrary(ctx, bookId)
	if err != nil {
		return 0, err
	}
	if !inLib { //случилась ситуация, когда хотим выдать пользователю книгу, которая не в библиотеке
		return 0, ErrBookNotInLibrary
	}

	setBookNotInLibraryQuery := fmt.Sprintf("UPDATE %s SET in_library = false WHERE id = $1", postgres.BooksTable)
	_, err = tx.Exec(ctx, setBookNotInLibraryQuery, bookId)
	if err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id, book_id) values ($1, $2) RETURNING id", postgres.UsersBooksTable)

	row := tx.QueryRow(ctx, query, userId, bookId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback(ctx)
		return 0, err
	}

	return id, tx.Commit(ctx)
}

func (b *BookRepo) DeleteBookFromUser(ctx context.Context, userId, bookId int) error {
	tx, err := b.db.Begin(ctx)
	if err != nil {
		return err
	}

	inLib, err := b.isBookInLibrary(ctx, bookId)
	if err != nil {
		return err
	}
	if inLib {
		return ErrBookInLibrary
	}

	setBookInLibraryQuery := fmt.Sprintf("UPDATE %s SET in_library = true WHERE id = $1", postgres.BooksTable)
	_, err = tx.Exec(ctx, setBookInLibraryQuery, bookId)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND book_id = $2", postgres.UsersBooksTable)
	res, err := tx.Exec(ctx, query, userId, bookId)

	if err != nil {
		tx.Rollback(ctx)
		return err
	}
	rowsAffected := res.RowsAffected() //случай, когда книга существует, не в библиотеке, но не числится у пользователя(мб. у другого)
	if rowsAffected == 0 {
		tx.Rollback(ctx)
		return ErrUserDoesNotHaveBook
	}

	return tx.Commit(ctx)
}

func (b *BookRepo) GetBooksByUser(ctx context.Context, userId int) ([]entity.Book, error) {
	//получаем id книг из id пользователя
	var books_id []int
	idQuery := fmt.Sprintf("SELECT book_id FROM %s WHERE user_id = $1", postgres.UsersBooksTable)

	rows, err := b.db.Query(ctx, idQuery, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		books_id = append(books_id, id)
	}

	//получаем книги по их id
	var books []entity.Book

	for _, id := range books_id {
		var book entity.Book
		bookQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", postgres.BooksTable)

		row := b.db.QueryRow(ctx, bookQuery, id)
		if err := row.Scan(&book.Id, &book.Title, &book.Author, &book.InLibrary); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
