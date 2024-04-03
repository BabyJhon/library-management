package repo

import "errors"

var (
	ErrUserNotFound        = errors.New("user not found")
	ErrBookNotFound        = errors.New("book not found")
	ErrBookInLibrary       = errors.New("the book is already in library")
	ErrBookNotInLibrary    = errors.New("the book is not in the library")
	ErrUserDoesNotHaveBook = errors.New("the user does not have this book")
)
