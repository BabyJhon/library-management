package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

// имена таблиц для sql запросов
const (
	UsersTable      = "users"
	BooksTable      = "books"
	UsersBooksTable = "users_books"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	db, openErr := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if openErr != nil {
		return nil, openErr
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}

	return db, nil
}
